# aws-order
亚马逊订单同步-golang


订单报告
http://docs.developer.amazonservices.com/en_US/reports/Reports_ReportType.html#ReportTypeCategories__OrderReports

获取订单
http://docs.developer.amazonservices.com/en_US/orders-2013-09-01/Orders_GetOrder.html


注册使用Amazon MWS
http://docs.developer.amazonservices.com/en_US/dev_guide/DG_Registering.html#DG_Registering

Amazon Marketplace Web Service（Amazon MWS）是一个使用签名进行身份验证的安全环境，并允许卖方使用Amazon MWS授权服务将调用权委派给开发人员。要有资格开发Amazon MWS 应用程序，您必须：

遵守可接受使用政策和数据保护政策。
有一个专业的销售计划。有关更多信息，请参阅卖方中心帮助中的销售计划。
提交Amazon MWS 开发人员注册和评估表。我们将评估您的提交并根据您所需的Amazon MWS访问级别为您分配角色。我们将要求您在注册时提交表格。
有关该表格的更多信息，请参见常见问题。

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

重要信息：您的秘密密钥是只有您和亚马逊才应该知道的秘密。请务必对其保密，以保护您的帐户。切勿在对亚马逊MWS的请求中包括您的密钥，也不要将其通过电子邮件发送给任何人。切勿与任何人共享它，即使他们声称来自亚马逊。亚马逊没有人会问您您的密钥。
重置您的AWS Access Key ID和密钥
如果您需要更改您的AWS Access Key ID和密钥，请通过聊天，电话或联系Amazon MWS表格与您本地市场的卖方支持联系，并要求重置您的AWS Access Key ID和密钥。如果您使用表单，请提供以下信息：

对于主题，输入“将访问密钥重置为我的Amazon MWS账户”。
对于您的评论，输入“请为我的账户重置AWS访问密钥ID和秘密密钥。”。
您将收到卖方支持团队的回复，其中包含指向卖方中心上“用户权限”页面的链接，您可以在其中为您的帐户生成新的密钥对。只有管​​理员可以使用该链接创建新密钥。

注意：当卖方支持部门删除您的密钥对时，所有使用当前密钥的应用程序都将被拒绝访问您的帐户信息，直到您用新密钥更新这些应用程序为止。
授权开发人员
卖方必须授权开发人员代表卖方致电Amazon MWS。

卖方授权开发商

开发人员必须是注册的亚马逊MWS 开发人员。请参阅注册为开发人员。
开发人员将其开发人员ID提供给卖方。
卖方进入卖方中心的“管理您的应用”页面，并以主要帐户持有人身份登录其亚马逊卖方帐户。
卖方单击“授权新开发者”按钮并遵循授权工作流程。
