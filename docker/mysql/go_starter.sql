CREATE TABLE `product` (
  `id` bigint(20) unsigned NOT NULL COMMENT '产品主键（用全局发号器）',
  `company_id` bigint(20) unsigned DEFAULT NULL,
  `title` varchar(256) NOT NULL COMMENT '产品标题',
  `intro` varchar(8192) NOT NULL DEFAULT '' COMMENT '产品简介',
  `desc_doc_id` bigint(20) unsigned NOT NULL COMMENT '关联的富文本文档',
  `artworks` varchar(1024) NOT NULL DEFAULT '' COMMENT '产品图片列表（JSON 串）',
  `is_interactive` tinyint(1) NOT NULL DEFAULT '0' COMMENT '能否在线交互',
  `is_public` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否公开（上架）',
  `public_at` timestamp NULL DEFAULT NULL COMMENT '公开时间',
  `extra` varchar(4096) NOT NULL DEFAULT '' COMMENT '额外的一些信息',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_interactive_prod` (`is_interactive`,`is_public`),
  KEY `idx_public_at` (`public_at`),
  KEY `idx_created_at` (`created_at`),
  KEY `idx_updated_at` (`updated_at`),
  KEY `idx_company` (`company_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='通用产品表（存放算法、硬件、解决方案等产品）';

CREATE TABLE `doc` (
	`id` bigint(20) unsigned NOT NULL COMMENT '文档 ID',
	`content_html` text COMMENT '渲染成 HTML 格式的描述信息',
	`content_md` text COMMENT 'Markdown 格式的文档（供后台编辑使用）',
	`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT 'Markdown 格式的文档，用于展示商品详情等';

CREATE TABLE `company` (
  `id` bigint(20) unsigned NOT NULL COMMENT '公司 ID',
  `title` varchar(64) NOT NULL DEFAULT '' COMMENT '公司名称',
  `intro` varchar(256) NOT NULL DEFAULT '' COMMENT '公司简介',
  `artworks` varchar(1024) NOT NULL DEFAULT '' COMMENT '公司图片',
  `url` varchar(128) NOT NULL DEFAULT '' COMMENT '官网链接',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='公司信息';

INSERT INTO `company` (`id`, `title`, `intro`, `artworks`, `url`, `created_at`, `updated_at`)
VALUES
	(34254759485505536, '知乎', '可信赖的问答社区，以让每个人高效获得可信赖的解答为使命。', '[\"v2-f8f2d20ee4cd5a1cc87c4ddba58f0078.jpg\"]', 'https://www.zhihu.com', '2020-01-16 11:30:11', '2020-01-16 11:30:11'),
	(34254759699415040, 'AIZOO', 'AI 算法商城，还有算法乐园等你体验', '[\"v2-f8f2d20ee4cd5a1cc87c4ddba581223444.jpg\"]', 'https://aizoo.com', '2020-01-16 11:30:11', '2020-01-16 11:30:11');

INSERT INTO `doc` (`id`, `content_html`, `content_md`, `created_at`, `updated_at`)
VALUES
	(123, 'hello, world\n', 'hello, world', '2020-01-16 12:29:44', '2020-01-16 12:29:44');

INSERT INTO `product` (`id`, `company_id`, `title`, `intro`, `desc_doc_id`, `artworks`, `is_interactive`, `is_public`, `public_at`, `extra`, `created_at`, `updated_at`)
VALUES
	(123, 34254759485505536, '测试', '测试', 123, '[\"v2-f8f2d20ee4cd5a1cc87c4ddba58f0078.jpg\"]', 1, 0, NULL, '', '2020-01-16 12:29:21', '2020-01-16 13:08:15'),
	(124, 34254759485505536, '测试', '测试', 123, '', 1, 0, NULL, '', '2020-01-16 12:29:21', '2020-01-16 13:08:17');
