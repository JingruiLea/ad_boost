# %%
import pandas as pd
from sqlalchemy import create_engine, text
import urllib.parse

# 数据库配置
DB_HOST = "120.26.166.254"
DB_PORT = 30006
DB_NAME = "taimer"
DB_USER = "taimer"
DB_PASSWORD = "Ta1mE!@#"

# URL 编码密码
encoded_password = urllib.parse.quote(DB_PASSWORD)

# 创建数据库引擎
engine = create_engine(
    f"mysql+pymysql://{DB_USER}:{encoded_password}@{DB_HOST}:{DB_PORT}/{DB_NAME}"
)
account_id = 1777719708674115
# SQL 查询
sql = f"SELECT * FROM ad_report_item WHERE advertiser_id = {account_id}"

# 使用 engine.connect() 并将 SQL 字符串转换为 text 对象
with engine.connect() as connection:
    df = pd.read_sql(text(sql), con=connection)

# 显示前几行数据以确认加载成功
print(df.head())


# %%
# 假设 df 是您的数据集
# df = ...

# 按消耗（StatCost）降序排序并去除重复的ad_id
top_ads_by_spend = (
    df.sort_values(by="stat_cost", ascending=False).drop_duplicates("ad_id").head()
)

# 按直接支付ROI（PrepayAndPayOrderRoi）降序排序并去除重复的ad_id
top_ads_by_roi = (
    df.sort_values(by="prepay_and_pay_order_roi", ascending=False)
    .drop_duplicates("ad_id")
    .head()
)

# 打印结果
print("Top Unique Ads by Spend:")
print(top_ads_by_spend[["ad_id", "ad_name", "stat_cost"]])
print("\nTop Unique Ads by ROI:")
print(top_ads_by_roi[["ad_id", "ad_name", "prepay_and_pay_order_roi"]])

# %%
# CREATE TABLE ad_report_item (
#     id bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
#     ad_id bigint(20) NOT NULL DEFAULT '0' COMMENT '广告计划id',
#     room_id bigint(20) NOT NULL DEFAULT '0' COMMENT 'roomid',
#     ad_name varchar(255) NOT NULL DEFAULT '' COMMENT '广告名字',
#     advertiser_id bigint(20) NOT NULL DEFAULT '0' COMMENT '广告主id',
#     click_cnt int(11) NOT NULL DEFAULT '0' COMMENT '点击次数',
#     convert_cnt int(11) NOT NULL DEFAULT '0' COMMENT '转化数',
#     convert_cost DECIMAL(10,2) NOT NULL DEFAULT '0' COMMENT '转化成本',
#     convert_rate DECIMAL(10,2) NOT NULL DEFAULT '0' COMMENT '转化率',
#     cpm_platform DECIMAL(10,2) NOT NULL DEFAULT '0' COMMENT '平均千次展示费用',
#     ctr DECIMAL(10,2) NOT NULL DEFAULT '0' COMMENT '点击率',
#     dy_follow int(11) NOT NULL DEFAULT '0' COMMENT '新增粉丝数',
#     pay_order_amount DECIMAL(10,2) NOT NULL DEFAULT '0' COMMENT '直接成交金额',
#     pay_order_count int(11) NOT NULL DEFAULT '0' COMMENT '直接成交订单数',
#     prepay_and_pay_order_roi DECIMAL(10,2) NOT NULL DEFAULT '0' COMMENT '直接支付roi',
#     show_cnt int(11) NOT NULL DEFAULT '0' COMMENT '展示次数',
#     stat_cost DECIMAL(10,2) NOT NULL DEFAULT '0' COMMENT '消耗',
#     roi_goal DECIMAL(10,2) NOT NULL DEFAULT '0' COMMENT 'roi',
#     cpa_bid DECIMAL(10,2) NOT NULL DEFAULT '0' COMMENT '出价',

#     created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
#     updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
#     deleted_at DATETIME(3) NULL DEFAULT NULL COMMENT '删除时间'

# ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='广告报表';

# %%
# 假设 df 是您的数据集
# df = ...
import matplotlib.pyplot as plt

# 设置特定的广告 ID
specific_ad_id = 1787781519916147

# 筛选特定广告 ID 的数据
specific_ad_data = df[df["ad_id"] == specific_ad_id]

# 确保时间戳是 datetime 类型
specific_ad_data["created_at"] = pd.to_datetime(specific_ad_data["created_at"])
specific_ad_data["updated_at"] = pd.to_datetime(specific_ad_data["updated_at"])

# 获取第一条和最后一条记录的时间
first_record = specific_ad_data["created_at"].min()
last_record = specific_ad_data["updated_at"].max()

# 打印结果
print("First record created at for ad_id", specific_ad_id, ":", first_record)
print("Last record updated at for ad_id", specific_ad_id, ":", last_record)
# %%
import matplotlib.pyplot as plt
import pandas as pd

# 假设 specific_ad_data 是您已经过滤后的特定广告ID的数据集
# specific_ad_data = ...

# 确保时间字段是 datetime 类型
specific_ad_data["created_at"] = pd.to_datetime(specific_ad_data["created_at"])

# 过滤掉 stat_cost 小于 100 的数据
filtered_data = specific_ad_data[specific_ad_data["stat_cost"] >= 100]

# 1. 散点图：Stat Cost vs Click Count
plt.figure(figsize=(10, 6))
plt.scatter(
    filtered_data["created_at"],
    filtered_data["click_cnt"],
    alpha=0.5,
    label="Click Count",
)
plt.xlabel("Time")
plt.ylabel("Count")
plt.title("Click Over Time")
plt.legend()
plt.show()

plt.scatter(
    filtered_data["created_at"],
    filtered_data["convert_cnt"],
    alpha=0.5,
    label="Convert Count",
)
plt.xlabel("Time")
plt.ylabel("Count")
plt.title("Convert Count Over Time")
plt.legend()
plt.show()

# 2. 散点图：Stat Cost vs ROI
plt.figure(figsize=(10, 6))
plt.scatter(
    filtered_data["created_at"],
    filtered_data["prepay_and_pay_order_roi"],
    alpha=0.5,
    label="ROI",
)
plt.xlabel("Time")
plt.ylabel("ROI")
plt.title("ROI Over Time")
plt.legend()
plt.show()

# 3. 散点图：Stat Cost vs Conversion Rate
plt.figure(figsize=(10, 6))
plt.scatter(
    filtered_data["created_at"],
    filtered_data["convert_rate"],
    alpha=0.5,
    label="Conversion Rate",
)
plt.xlabel("Time")
plt.ylabel("Conversion Rate")
plt.title("Conversion Rate Over Time")
plt.legend()
plt.show()

# 4. 散点图：Stat Cost vs CPM
plt.figure(figsize=(10, 6))
plt.scatter(
    filtered_data["created_at"], filtered_data["cpm_platform"], alpha=0.5, label="CPM"
)
plt.xlabel("Time")
plt.ylabel("CPM")
plt.title("CPM Over Time")
plt.legend()
plt.show()

# %%
# 假设 df 是您的数据集
# df = ...

# 分组并计算每个 ad_id 的采样数
ad_id_counts = df["ad_id"].value_counts()

# 找到采样最多的 ad_id
max_samples_ad_id = ad_id_counts.idxmax()
max_samples_count = ad_id_counts.max()

# 打印结果
print("Ad ID with the most samples:", max_samples_ad_id)
print("Number of samples for this Ad ID:", max_samples_count)


# %%
