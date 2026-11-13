## protobuf 协议定义

TODO: 是否分目录存储：csmsg, ssmsg, db, object

### Client <--> Server消息定义
- 文件名以cs_作为前缀，如cs_lobby.proto
- package命名：msg
- go_package命名：pbm

### Server <--> Server消息
- 文件名以ss_作为前缀，如ss_lobby.proto
- package命名：msg
- go_package命名：pbm

### 数据库存储对象定义
- 文件名以db_作为前缀，如db_account.proto
- package命名：db
- go_package命名：pbd

### 通用对象定义
- 文件名不加额外前缀
- package命名：obj
- go_package命名：pbo

## 命名规范
- 所有文件名采用小写字母，单词之间用下划线分隔，如：cs_lobby.proto，s_group_room.proto, cs_match_making.proto
- 所有字段名使用小写下划线风格（转成go语言的时候会自动改成驼峰命名法），如user_idd, room_id
- 所有枚举类型采用驼峰命名法，首字母大写，并加前缀E。如：EMatchMakingType
- 所有枚举值采用全大写下划线隔开风格。如：SERVICE_ID_LOGIN
  每个枚举类型的零值以NONE作为后缀。
- 所有消息名采用驼峰命名法，首字母大写。
  1. CS消息命名由四部分组成：前缀(CS) + 服务名(如Lobby， Account) + 功能定义(如Heartbeat， Login) + 消息类型(Req，Res，Ntf)。
     举例：CSLobbyHeartbeatReq，CSLobbyHeartbeatRes，CSLobbyJoinMatchNtf

     TODO: 是否去掉CS前缀？因为package名称已经明确是csmsg了
  2. SS消息命名由四部分组成：前缀(SS) + 服务名(如GroupRoom， MatchMaking) + 功能定义(如CreateRoom， JoinMatch) + 消息类型(Req，Res，Ntf)。
     举例：SSGroupRoomCreateRoomReq，SSGroupRoomCreateRoomRes，SSMatchMakingJoinMatchNtf

     TODO: 是否去掉SS前缀？因为package名称已经明确是ssmsg了
