#agent for pandora

## 简介
agent 监控当前机器上部署的各个 influxdb 的健康状态和相应挂载点的容量状况，并通过多个心跳将各个 influxdb 的监控指标发送给 scheduler。注意，如果某个 influxdb 不可读，则不会发送对应该 influxdb 的心跳。
	ID          string  `json:"diskTag"`     // database id	Host        string  `json:"host"`        // host
	Port        string  `json:"port"`        // port
	SpaceUsed   uint64  `json:"spaceUsed"`   // used bytes
	SpaceFree   uint64  `json:"spaceFree"`   // free bytes
	ConsumeRate float64 `json:"consumeRate"` // consumed bytes per second
	Readable    bool    `json:"readable"`    // 是否可读
	Writeable   bool    `json:"writeable"`   //是否可写

## 要求
1. 不同的 influxdb 必须使用不同的挂载点
2. 需要在 agent 的配置文件中给出各个 influxdb 的配置文件的绝对路径
3. 每个 influxdb 中都有一个 agent 专用的 database，名称为 com_qiniu_pandora_agent_ping

## influxdb 监控指标
1. 占用的 HTTP 端口号 （对应 influxdb 配置文件中 http 小节中的 bind-address 字段）
2. 是否可写
3. 是否可读

## influxdb 挂载点监控指标
1. 挂载点已用容量 （单位：字节）
2. 挂载点可用容量 （单位：字节）
3. 挂载点容量消耗速度 （单位：字节/秒）
        - 计算公式 = (当前可用容量字节数 - 上一次心跳中的可用容量字节数) / 心跳间隔秒数
        - 正数表示可用容量在减小
        - 负数表示可用容量在增加 （例如，当 influxdb 清理 WAL 日志时，可用容量就会增加）

## 如何监控 influxdb
1. 检查其是否可写
2. 检查其是否可读

## 如何监控 influxdb 挂载点
1.  

