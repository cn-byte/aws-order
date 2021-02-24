# aws-order
亚马逊订单同步-golang


订单报告
http://docs.developer.amazonservices.com/en_US/reports/Reports_ReportType.html#ReportTypeCategories__OrderReports

获取订单
http://docs.developer.amazonservices.com/en_US/orders-2013-09-01/Orders_GetOrder.html


注册使用Amazon MWS
http://docs.developer.amazonservices.com/en_US/dev_guide/DG_Registering.html#DG_Registering

注册为开发者
转到卖方中心中的“用户权限”页面，并以主要帐户持有人身份登录到您的亚马逊销售帐户。
在Amazon MWS开发人员访问密钥下，单击 访问开发人员凭证按钮。
在“开发人员中心”页面上的“您是软件开发人员在使用MWS构建应用程序吗？”下。，单击此处申请访问链接，然后按照Amazon MWS 开发人员注册和评估表单上的说明进行操作。
我们将评估您在表单中提供的信息，然后跟踪支持案例日志和后续步骤。有关开发者注册和评估的更多信息，请参阅常见问题。

使用您的Amazon MWS开发人员ID和访问密钥
重要： 不要共享Amazon MWS 访问密钥
共享您的亚马逊MWS访问密钥违反了亚马逊的服务条款。共享访问密钥可能会导致您的销售帐户和访问密钥被暂停。

以下是Amazon MWS 开发人员ID和访问密钥的示例：

开发者ID（12位数字的标识符）： 1234-3214-4321
AWS Access Key ID（20个字符的字母数字标识符）： 022QF0EXAMPLEH9DHM02
密钥（40个字符的标识符）： kWcrlEXAMPLEM / LtmEENI / aVmYvHNif5zB + d9 + ct
如果您正在为卖家开发Amazon MWS应用程序或向卖家提供与Amazon MWS相关的开发服务，则必须向这些卖家提供开发人员ID，以便他们可以授权您使用Amazon MWS访问其卖家账户。要求卖方拥有自己的Amazon MWS开发人员凭证是违反Amazon服务条款的行为。

AWS Access Key ID与您的Amazon MWS注册相关联。您将其包括在所有Amazon MWS请求中，以将自己标识为请求的发送者。AWS Access Key ID不是秘密。为了提供证明您确实是请求的发送者，您还必须包括一个数字签名。对于除使用Amazon MWS生成的请求以外的所有请求客户端库，您可以使用密钥来计算签名。Amazon在请求中使用AWS Access Key ID查找您的秘密密钥，然后使用该密钥计算数字签名。如果Amazon计算的签名与您发送的签名匹配，则该请求被视为真实请求。否则，请求将无法通过身份验证，并且不会被处理。
