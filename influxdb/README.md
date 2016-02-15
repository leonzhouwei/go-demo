#agent for pandora

## 简介
agent 监控当前机器上部署的各个 influxdb 及其所在的挂载点，每隔一定秒数（可通过 agent 的配置文件指定）更新一次各个 influxdb 的健康状态和各个挂载点的容量状况，并通过多个心跳将更新后的各个监控指标发送给 scheduler。注意，如果某个 influxdb 不可读或者其挂载点出错，则不会发送对应该 influxdb 的心跳。

## 要求
1. 不同的 influxdb 必须处于不同的挂载点下
2. 需要在 agent 的配置文件中给出各个 influxdb 的配置文件的绝对路径
3. 每个 influxdb 中都有一个 agent 专用的 database，名称为 com_qiniu_pandora_agent_ping

## influxdb 监控指标
1. 是否可写
3. 是否可读

## influxdb 挂载点监控指标
- 挂载点已用容量 （单位：字节）
- 挂载点可用容量 （单位：字节）
- 挂载点容量消耗速度 （单位：字节/秒）
        - 计算公式 = (当前可用容量字节数 - 上一次心跳中的可用容量字节数) / 心跳间隔秒数
        - 正数表示可用容量在减小
        - 负数表示可用容量在增加 （例如，当 influxdb 清理 WAL 日志时，可用容量就会增加）

## 如何监控 influxdb
1. 是否可读：通过 HTTP API 向 agent 专用的 com_qiniu_pandora_agent_ping 写入 cpu,host=s1 value=10 1，如果出错，则该 influxdb 不可写
2. 是否可读：通过 HTTP API 对 agent 专用的 com_qiniu_pandora_agent_ping 执行 select * from cpu 操作，如果出错，则该 influxdb 不可读

## 如何监控 influxdb 挂载点
和 [influxdata/telegraf](https://github.com/influxdata/telegraf) 一样，使用 [shirou/gopsutil](https://github.com/shirou/gopsutil) 的 disk 包来获取挂载点的磁盘使用状况

## 心跳的数据结构
1. diskTag (influxdb 实例使用的挂载点的 ID)
2. port (influxdb 的 HTTP 端口号，对应 influxdb 配置文件中 http 小节内的 bind-address 字段)
3. readable (influxdb 是否可读，可读为 true，不可读为 false)
4. writeable (influxdb 是否可写，可写为 true，不可写为 false)
5. spaceUsed (挂载点已用容量)
6. spaceFree (挂载点可用容量)
7. consumeRate (挂载点容量消耗速度)

