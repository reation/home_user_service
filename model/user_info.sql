CREATE TABLE `user_info`  (
                              `id` int(11) NOT NULL AUTO_INCREMENT,
                              `login_username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '登录用户名',
                              `login_password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '登录密码',
                              `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '昵名',
                              `icon` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '头像地址',
                              `gender` tinyint(4) NOT NULL DEFAULT 0 COMMENT '性别 0：男，1：女',
                              `status` tinyint(4) NOT NULL DEFAULT 1 COMMENT '账户状态 1：正常，2：冻结，3：取消，4：异常',
                              `type` tinyint(4) NOT NULL DEFAULT 0 COMMENT '账户类型 1：买家，2：卖家，3：工作人员',
                              `b_id` int(11) NOT NULL DEFAULT 0 COMMENT '所属商家',
                              `update_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
                              `create_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
                              PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci;