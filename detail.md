**不关心具体库的对应 Demo 可以跳转至文末看三者的对比。**

原本我们在 `Go` 中使用的是 [gqlgen](https://github.com/99designs/gqlgen) ，这个库通过生成代码的方式保证类型安全和类型绑定，以减少我们在写代码时的样板代码。大部分情况下使用起来很不错，但是 `gqlgen` 每次更改 `schema.graphql` 后都需要重新生成代码，这就需要我们不能修改生成的代码，而要将每次生成的 `resolver` 拷贝出来进行使用，为了便于管理还需要组织进不同的文件内（由于需要循环引用，所以不能分在多个 `package` 中）。考虑到以后可能将越来越多的接口迁移至 `Go` 中，这样的代码可能不太适合以后的项目结构，所需准备尝试新的 `GraphQL` 库。

下面是 `gqlgen` 官网的[特性对比](https://gqlgen.com/feature-comparison/)：

本文采用以下简单 `Demo` 分别用三种库分别实现了 [schema.graphql](./schema.graphql) 对应的代码。

#### [gqlgen](https://github.com/99designs/gqlgen)
接触 `Go` 以来感觉到它的哲学就是生成代码， `gqlgen` 也是符合这一哲学的一员。

我们直接参照官方的配置文件配置即可，为了更好的组织代码结构，我们将生成的代码都放入 `gen` 文件夹下。

配置文件 [gqlgen.yml](./gqlgen/gqlgen.yml) 中：
- `models.Todo.model` 可以实现类型绑定，即：不需要生成对应的类型 ，直接使用我们已存在的实体。
- `models.Todo.fields.id` 可以实现自定义 `resolver` ，即：对某个对象的某个字段需要特殊处理，不直接返回对应的值。

运行完 `go run github.com/99designs/gqlgen -v` 后就会生成对应的三个文件：
- [generated.go](./gqlgen/gen/generated.go) 中包含了核心逻辑，主要用来驱动整个 `GraphQL` 运转，这部分的代码开发时不会使用到。
- [objects_gen.go](./gqlgen/gen/objects_gen.go) 中包含了 `schema.graphql` 中定义的类型： `ToggleTodoOutput` 和 `CreateTodoInput` ，而 `Todo` 将与 `models.Todo` 绑定，不进行生成。
- [resolver_gen.go](./gqlgen/gen/resolver_gen.go) 中包含了所有自定义的 `resolver` 的签名。由于接口会不断增长，所以我们不直接改写这个文件，而是将其中的内容拷贝出来，并在组织好的文件中进行使用。这样以后新增接口时，只需将增加的 `resolver` 拷贝至对应的文件中实现即可。

在当前比较简单清晰的情况下，我们将使用四个文件完成所需的全部业务，后续随着接口增多，可以按照业务细化为更多的文件（由于需要循环引用，所以所有的文件只能在同一个 `package` 中）：

- [mutations.go](./gqlgen/mutations.go) 中为所有 `Mutation` 的实现逻辑，在这里处理对应接口的核心逻辑

- [queries.go](./gqlgen/queries.go) 中为所有 `Query` 的实现逻辑，在这里处理对应接口的核心逻辑

- [resolvers.go](./gqlgen/resolvers.go) 中为所有的自定义的 `resolver` 实现逻辑，在这里对每个字段返回前进行一些自定义的处理，然后返回所需的值（本 `Demo` 中，我们让所有的 `id` 都比真实值大 `1`）

- [schema.go](./gqlgen/schema.go) 中定义了处理接口所需的入口 —— `Schema` ，以及对应的顶层 `resolver` ，包含 `Query` 和 `Mutation`

经过以上 `Demo` 的使用可以发现： `gqlgen` 非常方便，能自动帮我们生成大部分代码，我们只需关注我们自己的核心业务逻辑和对应字段的解析实现即可，其他都会使用默认的代码。

总结一下我在实现本 `Demo` 时发现的一些优缺点：
- 几乎没有样板代码，代码量少（仅需 `122` 行代码（包括 `.yml` 配置文件）即可完成 `Demo`），可以快速进行开发
- 参数不支持 `defaultValue` ，尽管我们已经在 `schema.graphql` 中定义过 `CreateTodoInput.completed` 的默认值，但是在运行时不能自动注入默认值，需要手动处理，在有很多默认值的情况下可能会在业务代码中产生很多样板代码，而且还需要时刻保持所有地方的默认值一致，以免前端看到的和实际不一样
- 前端获取到的接口信息不含注释，尽管我们已经在 `schema.graphql` 中详细定义过所有注释，但是这些信息无法在前端显示，可能会增加沟通成本

以上 `Demo` 及后续 `graphql-go` 和 `gophers` 的 `Demo` 可以在 [go-graphql-comparison](https://github.com/idealism-xxm/go-graphql-comparison) 中找到。

在最初选择库的时候，我其实比较推崇 [graphql-go](https://github.com/graphql-go/graphql) ，因为它和 `Graphene` 使用的 `GraphQL` 库都参照了 [graphql-js](https://github.com/graphql/graphql-js) 。它们使用方式类似，且拥有很多可配置的字段，所以我认为切换到这个库的使用成本比较小。但由于时间紧急且需要迁移的接口比较多，就先选择了 `gqlgen` 。

#### [graphql-go](https://github.com/graphql-go/graphql)
本 Demo 很简单，所以没有分成多个 `package` ，定义的类型和变量也基本全以小写开始。当项目变得复杂时，可以按照业务分成不同的 `package` ，不过还是推荐大部分以小写开始，因为出除了需要对外暴露的变量和类型，其余都只会在本 `package` 内使用。

- [inputs.go](./graphqlgo/inputs.go) 中定义了入参类型

- [objects.go](./graphqlgo/objects.go) 中定义了可解析的类型，一般和 `model` 对应

- [outputs.go](./graphqlgo/outputs.go) 中定义了 `Mutation` 返回的类型，主要是继承了 `Graphene` 中的习惯，原来我们会在其中加入 `ok` 和 `error` 字段用以返回额外信息，后来渐渐替换为使用 `GraphQL` 自带的 `errors` 。
    - 由于 `graphql-go` 是动态类型，所以其中需要定义了返回值类型配置后，再定义其对应的 `struct` 便于后续 `resolver` 

- [mutations.go](./graphqlgo/mutations.go) 中定义了所有的 `Mutation` 和顶层的 `Mutation` 的声明

- [queries.go](./graphqlgo/queries.go) 中定义了所有的 `Query` 和顶层的 `Query` 的声明

- [schema.go](./graphqlgo/queries.go) 中定义了 `Schema` ，用于在启动时使用。基本只需要暴露这个变量即可

经过以上 Demo 的使用可以发现： `graphql-go` 虽然在思维上不需要转换太多，但是由于相应的方式在静态语言中很难优雅处理，所以处处都是样板代码，且处处需要类型转换，并且业务代码都隐藏在复杂的配置中，很难快速找到想要的逻辑。刚开始我还充满了信心，想着类型可以通过反射和 `Tag` 用方法进行返回（但还是无法避免侵入业务），后来就直接放弃了，因为真正的业务逻辑还不到需要的代码的一半，会极大降低开发效率，且一样难以维护。总共需要 `234` 行完成 Demo 。

最后总结一下我在用 `graphql-go` 实现本 Demo 时发现的一些优缺点：

- 可配置项全面，前端获取到的接口信息含有注释
- 处处都是样板代码，处处需要类型转换
- 代码量大，真实业务逻辑占比不到一半，且隐藏在复杂的配置中


#### [gophers](https://github.com/graph-gophers/graphql-go)
最后就尝试了 `gophers` ，看起来和 `gqlgen` 实现方式类似，不过不需要生成代码，并且基本也只需要关注真正的业务逻辑，应该是 `gqlgen` 的强力竞争者。

`gophers` 直接使用 `schema.graphql` 生成所需的信息，所以不需要在 `objects` 定义类型，而是通过 `resolver` 的方式解析每个类型的字段，并且只有入参的类型需要定义，其他都只需要 `resolver`。对于基本类型，我们可以直接返回对应的类型即可；对于非基本类型，我们就只需要返回对应的 `resolver` 类型。我们可以将可提供所需字段的变量成为 `resolver` 类型的私有变量，既可以只在内部使用，有又不会与对应的解析函数冲突。由于 `gophers` 的限制，我们的 `resolver` 函数和入参的字段必须公开，只有对应的类型可以私有。

- [inputs.go](./gophers/inputs.go) 中定义了入参类型，可以发现和 `gqlgen` 生成的基本一致。不过有默认值的非必传字段不是引用，它会在不传时使用默认值，可以减少业务代码的判断

- [outputs.go](./gophers/outputs.go) 中定义了 `Mutation` 返回的 `resolver` 类型

- [resolvers.go](./gophers/resolvers.go)  中定义了可输出解析的类型的 `resolver` 类型

- [mutations.go](./gophers/mutations.go) 中定义了所有的 `Mutation` 的 `resolver` 函数和顶层 `Mutation` 类型

- [queries.go](./gophers/queries.go) 中定义了所有的 `Query` 的 `resolver` 函数和顶层 `Query` 类型

- [schema.go](./gophers/schema.go) 中定义了 `Schema` ，用于在启动时使用。和 `graphql-go` 一样，基本只需要暴露这个变量即可

经过以上 Demo 的使用可以发现： `gophers` 的思路和 `gqlgen` 一样，只需要定义对应的 `resolver` 的类型即可，基本不存在样板代码。并且代码量和使用 `gqlgen` 时相当，仅需 `130` 行即可完成 Demo 。

最后总结一下我在用 `gophers` 实现本 Demo 时发现的一些优缺点：

- 代码量小，样板代码也少，可以支持快速开发
- 有默认值的入参会在不传时使用默认值，并且对应的类型不是指针，可以减少业务代码的判断
-返回非基本类型的列表时需要手动封装成对应的 `resolver` 类型的列表


再总结三者的对比：

| | gqlgen | graphql-go | gophers
---|---|---|---
**代码行数** | 122 | 234 | 130
**样板代码** | 少 | 多 | 少
**入参默认值** | ✕ | √ | √
**接口信息含有注释** | ✕ | √ | √
**非基本类型列表无需封装** | √ | √ | x
**每秒可处理请求数** | 49925 | 19004 | 44308

`每秒可处理请求数` 数据来源：[golang-graphql-benchmark](https://github.com/appleboy/golang-graphql-benchmark)

**注**：本对比仅代表当前 Demo 使用感受，增加了部分 [gqlgen 官方特性对比](https://gqlgen.com/feature-comparison/) 未体现的对比。
