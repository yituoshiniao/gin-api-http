/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80019
 Source Host           : localhost:3306
 Source Schema         : db_goods_center

 Target Server Type    : MySQL
 Target Server Version : 80019
 File Encoding         : 65001

 Date: 11/05/2024 17:12:23
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for territory_code
-- ----------------------------
DROP TABLE IF EXISTS `territory_code`;
CREATE TABLE `territory_code` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `number` bigint DEFAULT NULL COMMENT '全局唯一number',
  `cn_abbreviation` varchar(100) NOT NULL DEFAULT '' COMMENT '中文简称',
  `en_abbreviation` varchar(100) NOT NULL DEFAULT '' COMMENT '英文简称',
  `en_full_name` varchar(100) NOT NULL DEFAULT '' COMMENT '英文简称',
  `two_letter_code` varchar(10) NOT NULL DEFAULT '' COMMENT '两字母代码(大写)',
  `three_letter_code` varchar(10) NOT NULL DEFAULT '' COMMENT '三字母代码(大写)',
  `numeric_code` varchar(10) NOT NULL DEFAULT '' COMMENT '数字代码',
  `description` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '备注信息',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除: 1 是; 0 否;',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_two_letter_code` (`two_letter_code`),
  UNIQUE KEY `uk_three_letter_code` (`three_letter_code`),
  UNIQUE KEY `uk_number` (`number`)
) ENGINE=InnoDB AUTO_INCREMENT=246 DEFAULT CHARSET=utf8 COMMENT='国家地区码信息';

-- ----------------------------
-- Records of territory_code
-- ----------------------------
BEGIN;
INSERT INTO `territory_code` VALUES (1, 1604771561931374592, '阿富汗', 'Afghanistan', 'the Islamic Republic of Afghanistan', 'AF', 'AFG', '004', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (2, 1604771561931374593, '奥兰群岛', 'Aland Islands', '', 'AX', 'ALA', '248', 'ISO 3166-1:2006新增', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (3, 1604771561931374594, '阿尔巴尼亚', 'Albania', 'the Republic of Albania', 'AL', 'ALB', '008', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (4, 1604771561931374595, '阿尔及利亚', 'Algeria', 'the People\'s Democratic Republic of Algeria', 'DZ', 'DZA', '012', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (5, 1604771561931374596, '美属萨摩亚', 'American Samoa', '', 'AS', 'ASM', '016', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (6, 1604771561931374597, '安道尔', 'Andorra', 'the Principality of Andorra', 'AD', 'AND', '020', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (7, 1604771561931374598, '安哥拉', 'Angola', 'the Republic of Angola', 'AO', 'AGO', '024', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (8, 1604771561931374599, '安圭拉', 'Anguilla', '', 'AI', 'AIA', '660', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (9, 1604771561931374600, '南极洲', 'Antarctica', '', 'AQ', 'ATA', '010', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (10, 1604771561931374601, '安提瓜和巴布达', 'Antigua and Barbuda', '', 'AG', 'ATG', '028', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (11, 1604771561931374602, '阿根廷', 'Argentina', 'the Argentine Republic', 'AR', 'ARG', '032', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (12, 1604771561931374603, '亚美尼亚', 'Armenia', 'the Republic of Armenia', 'AM', 'ARM', '051', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (13, 1604771561931374604, '阿鲁巴', 'Aruba', '', 'AW', 'ABW', '533', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (14, 1604771561931374605, '澳大利亚', 'Australia', '', 'AU', 'AUS', '036', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (15, 1604771561931374606, '奥地利', 'Austria', 'the Republic of Austria', 'AT', 'AUT', '040', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (16, 1604771561931374607, '阿塞拜疆', 'Azerbaijan', 'the Republic of Azerbaijan', 'AZ', 'AZE', '031', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (17, 1604771561931374608, '巴哈马', 'Bahamas (The)', 'the Commonwealth of The Bahamas', 'BS', 'BHS', '044', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (18, 1604771561931374609, '巴林', 'Bahrain', 'the Kingdom of Bahrain', 'BH', 'BHR', '048', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (19, 1604771561931374610, '孟加拉国', 'Bangladesh', 'the People\'s Republic of Bangladesh', 'BD', 'BGD', '050', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (20, 1604771561931374611, '巴巴多斯', 'Barbados', '', 'BB', 'BRB', '052', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (21, 1604771561931374612, '白俄罗斯', 'Belarus', 'the Republic of Belarus', 'BY', 'BLR', '112', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (22, 1604771561931374613, '比利时', 'Belgium', 'the Kingdom of Belgium', 'BE', 'BEL', '056', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (23, 1604771561931374614, '伯利兹', 'Belize', '', 'BZ', 'BLZ', '084', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (24, 1604771561931374615, '贝宁', 'Benin', 'the Republic of Benin', 'BJ', 'BEN', '204', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (25, 1604771561931374616, '百慕大', 'Bermuda', '', 'BM', 'BMU', '060', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (26, 1604771561931374617, '不丹', 'Bhutan', 'the Kingdom of Bhutan', 'BT', 'BTN', '064', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (27, 1604771561931374618, '玻利维亚', 'Bolivia', 'the Republic of Bolivia', 'BO', 'BOL', '068', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (28, 1604771561931374619, '波黑', 'Bosnia and Herzegovina', '', 'BA', 'BIH', '070', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (29, 1604771561931374620, '博茨瓦纳', 'Botswana', 'the Republic of Botswana', 'BW', 'BWA', '072', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (30, 1604771561931374621, '布维岛', 'Bouvet Island', '', 'BV', 'BVT', '074', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (31, 1604771561931374622, '巴西', 'Brazil', 'the Federative Republic of Brazil', 'BR', 'BRA', '076', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (32, 1604771561931374623, '英属印度洋领地', 'British Indian Ocean Territory (the)', '', 'IO', 'IOT', '086', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (33, 1604771561931374624, '文莱', 'Brunei Darussalam', '', 'BN', 'BRN', '096', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (34, 1604771561931374625, '保加利亚', 'Bulgaria', 'the Republic of Bulgaria', 'BG', 'BGR', '100', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (35, 1604771561931374626, '布基纳法索', 'Burkina Faso', 'Burkina Faso', 'BF', 'BFA', '854', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (36, 1604771561931374627, '布隆迪', 'Burundi', 'the Republic of Burundi', 'BI', 'BDI', '108', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (37, 1604771561981706240, '柬埔寨', 'Cambodia', 'the Kingdom of Cambodia', 'KH', 'KHM', '116', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (38, 1604771561981706241, '喀麦隆', 'Cameroon', 'the Republic of Cameroon', 'CM', 'CMR', '120', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (39, 1604771561981706242, '加拿大', 'Canada', '', 'CA', 'CAN', '124', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (40, 1604771561981706243, '佛得角', 'Cape Verde', 'the Republic of Cape Verde', 'CV', 'CPV', '132', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (41, 1604771561981706244, '开曼群岛', 'Cayman Islands (the)', '', 'KY', 'CYM', '136', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (42, 1604771561981706245, '中非', 'Central African Republic (the)', 'the Central African Republic', 'CF', 'CAF', '140', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (43, 1604771561981706246, '乍得', 'Chad', 'the Republic of Chad', 'TD', 'TCD', '148', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (44, 1604771561981706247, '智利', 'Chile', 'the Republic of Chile', 'CL', 'CHL', '152', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (45, 1604771561981706248, '中国', 'China', 'the People\'s Republic of China', 'CN', 'CHN', '156', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (46, 1604771561981706249, '圣诞岛', 'Christmas Island', '', 'CX', 'CXR', '162', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (47, 1604771561981706250, '科科斯（基林）群岛', 'Cocos (Keeling) Islands (the)', '', 'CC', 'CCK', '166', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (48, 1604771561981706251, '哥伦比亚', 'Colombia', 'the Republic of Colombia', 'CO', 'COL', '170', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (49, 1604771561981706252, '科摩罗', 'Comoros', 'the Union of the Comoros', 'KM', 'COM', '174', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (50, 1604771561981706253, '刚果（布）', 'Congo', 'the Republic of the Congo', 'CG', 'COG', '178', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (51, 1604771561981706254, '刚果（金）', 'Congo (the Democratic Republic of the)', 'the Democratic Republic of the Congo', 'CD', 'COD', '180', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (52, 1604771561981706255, '库克群岛', 'Cook Islands (the)', '', 'CK', 'COK', '184', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (53, 1604771561981706256, '哥斯达黎加', 'Costa Rica', 'the Republic of Costa Rica', 'CR', 'CRI', '188', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (54, 1604771561981706257, '科特迪瓦', 'Côte d\'Ivoire', 'the Republic of Côte d\'Ivoire', 'CI', 'CIV', '384', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (55, 1604771561981706258, '克罗地亚', 'Croatia', 'the Republic of Croatia', 'HR', 'HRV', '191', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (56, 1604771561981706259, '古巴', 'Cuba', 'the Republic of Cuba', 'CU', 'CUB', '192', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (57, 1604771561981706260, '塞浦路斯', 'Cyprus', 'the Republic of Cyprus', 'CY', 'CYP', '196', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (58, 1604771561981706261, '捷克', 'Czech Republic (the)', 'the Czech Republic', 'CZ', 'CZE', '203', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (59, 1604771561981706262, '丹麦', 'Denmark', 'the Kingdom of Denmark', 'DK', 'DNK', '208', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (60, 1604771561981706263, '吉布提', 'Djibouti', 'the Republic of Djibouti', 'DJ', 'DJI', '262', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (61, 1604771561981706264, '多米尼克', 'Dominica', 'the Commonwealth of Dominica', 'DM', 'DMA', '212', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (62, 1604771561981706265, '多米尼加', 'Dominican Republic (the)', 'the Dominican Republic', 'DO', 'DOM', '214', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (63, 1604771561981706266, '厄瓜多尔', 'Ecuador', 'the Republic of Ecuador', 'EC', 'ECU', '218', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (64, 1604771561981706267, '埃及', 'Egypt', 'the Arab Republic of Egypt', 'EG', 'EGY', '818', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (65, 1604771561981706268, '萨尔瓦多', 'El Salvador', 'the Republic of El Salvador', 'SV', 'SLV', '222', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (66, 1604771561981706269, '赤道几内亚', 'Equatorial Guinea', 'the Republic of Equatorial Guinea', 'GQ', 'GNQ', '226', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (67, 1604771561981706270, '厄立特里亚', 'Eritrea', '', 'ER', 'ERI', '232', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (68, 1604771561981706271, '爱沙尼亚', 'Estonia', 'the Republic of Estonia', 'EE', 'EST', '233', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (69, 1604771561981706272, '埃塞俄比亚', 'Ethiopia', 'the Federal Democratic Republic of Ethiopia', 'ET', 'ETH', '231', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (70, 1604771561981706273, '福克兰群岛（马尔维纳斯）', 'Falkland Islands (the) [Malvinas]', '', 'FK', 'FLK', '238', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (71, 1604771561981706274, '法罗群岛', 'Faroe Islands (the)', '', 'FO', 'FRO', '234', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (72, 1604771561981706275, '斐济', 'Fiji', 'the Republic of the Fiji Islands', 'FJ', 'FJI', '242', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (73, 1604771561981706276, '芬兰', 'Finland', 'the Republic of Finland', 'FI', 'FIN', '246', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (74, 1604771561981706277, '法国', 'France', 'the French Republic', 'FR', 'FRA', '250', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (75, 1604771561981706278, '法属圭亚那', 'French Guiana', '', 'GF', 'GUF', '254', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (76, 1604771561981706279, '法属波利尼西亚', 'French Polynesia', '', 'PF', 'PYF', '258', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (77, 1604771561981706280, '法属南部领地', 'French Southern Territories (the)', '', 'TF', 'ATF', '260', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (78, 1604771561981706281, '加蓬', 'Gabon', 'the Gabonese Republic', 'GA', 'GAB', '266', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (79, 1604771561981706282, '冈比亚', 'Gambia (The)', 'the Republic of The Gambia', 'GM', 'GMB', '270', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (80, 1604771561981706283, '格鲁吉亚', 'Georgia', '', 'GE', 'GEO', '268', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (81, 1604771561981706284, '德国', 'Germany', 'he Federal Republic of Germany', 'DE', 'DEU', '276', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (82, 1604771561981706285, '加纳', 'Ghana', 'the Republic of Ghana', 'GH', 'GHA', '288', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (83, 1604771561981706286, '直布罗陀', 'Gibraltar', '', 'GI', 'GIB', '292', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (84, 1604771561981706287, '希腊', 'Greece', 'the Hellenic Republic', 'GR', 'GRC', '300', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (85, 1604771561981706288, '格陵兰', 'Greenland', '', 'GL', 'GRL', '304', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (86, 1604771561981706289, '格林纳达', 'Grenada', '', 'GD', 'GRD', '308', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (87, 1604771561981706290, '瓜德罗普', 'Guadeloupe', '', 'GP', 'GLP', '312', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (88, 1604771561981706291, '关岛', 'Guam', '', 'GU', 'GUM', '316', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (89, 1604771561981706292, '危地马拉', 'Guatemala', 'the Republic of Guatemala', 'GT', 'GTM', '320', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (90, 1604771561981706293, '格恩西岛', 'Guernsey', '', 'GG', 'GGY', '831', 'ISO 3166-1:2006新增', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (91, 1604771561981706294, '几内亚', 'Guinea', 'the Republic of Guinea', 'GN', 'GIN', '324', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (92, 1604771561981706295, '几内亚比绍', 'Guinea-Bissau', 'the Republic of Guinea-Bissau', 'GW', 'GNB', '624', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (93, 1604771561981706296, '圭亚那', 'Guyana', 'the Republic of Guyana', 'GY', 'GUY', '328', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (94, 1604771561985900544, '海地', 'Haiti', 'the Republic of Haiti', 'HT', 'HTI', '332', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (95, 1604771561985900545, '赫德岛和麦克唐纳岛', 'Heard Island and McDonald Islands', '', 'HM', 'HMD', '334', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (96, 1604771561985900546, '梵蒂冈', 'Holy See (the) [Vatican City State]', '', 'VA', 'VAT', '336', 'ISO 3166.1:2006调整英文名称，代码未变', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (97, 1604771561985900547, '洪都拉斯', 'Honduras', 'the Republic of Honduras', 'HN', 'HND', '340', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (98, 1604771561985900548, '香港', 'Hong Kong', 'the Hong Kong Special Administrative Region of China', 'HK', 'HKG', '344', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (99, 1604771561985900549, '匈牙利', 'Hungary', 'the Republic of Hungary', 'HU', 'HUN', '348', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (100, 1604771561985900550, '冰岛', 'Iceland', 'the Republic of Iceland', 'IS', 'ISL', '352', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (101, 1604771561985900551, '印度', 'India', 'the Republic of India', 'IN', 'IND', '356', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (102, 1604771561985900552, '印度尼西亚', 'Indonesia', 'the Republic of Indonesia', 'ID', 'IDN', '360', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (103, 1604771561985900553, '伊朗', 'Iran (the Islamic Republic of)', 'the Islamic Republic of Iran', 'IR', 'IRN', '364', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (104, 1604771561985900554, '伊拉克', 'Iraq', 'the Republic of Iraq', 'IQ', 'IRQ', '368', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (105, 1604771561985900555, '爱尔兰', 'Ireland', '', 'IE', 'IRL', '372', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (106, 1604771561985900556, '英国属地曼岛', 'Isle of Man', '', 'IM', 'IMN', '833', 'ISO 3166-1:2006新增', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (107, 1604771561985900557, '以色列', 'Israel', 'the State of Israel', 'IL', 'ISR', '376', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (108, 1604771561985900558, '意大利', 'Italy', 'the Republic of Italy', 'IT', 'ITA', '380', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (109, 1604771561985900559, '牙买加', 'Jamaica', '', 'JM', 'JAM', '388', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (110, 1604771561985900560, '日本', 'Japan', '', 'JP', 'JPN', '392', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (111, 1604771561985900561, '泽西岛', 'Jersey', '', 'JE', 'JEY', '832', 'ISO 3166-1:2006新增', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (112, 1604771561985900562, '约旦', 'Jordan', 'the Hashemite Kingdom of Jordan', 'JO', 'JOR', '400', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (113, 1604771561985900563, '哈萨克斯坦', 'Kazakhstan', 'the Republic of Kazakhstan', 'KZ', 'KAZ', '398', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (114, 1604771561985900564, '肯尼亚', 'Kenya', 'the Republic of Kenya', 'KE', 'KEN', '404', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (115, 1604771561985900565, '基里巴斯', 'Kiribati', 'the Republic of Kiribati', 'KI', 'KIR', '296', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (116, 1604771561985900566, '朝鲜', 'Korea (the Democratic People\'s Republic of)', 'the Democratic People\'s Republic of Korea', 'KP', 'PRK', '408', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (117, 1604771561985900567, '韩国', 'Korea (the Republic of)', 'the Republic of Korea', 'KR', 'KOR', '410', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (118, 1604771561985900568, '科威特', 'Kuwait', 'he State of Kuwait', 'KW', 'KWT', '414', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (119, 1604771561985900569, '吉尔吉斯斯坦', 'Kyrgyzstan', 'the Kyrgyz Republic', 'KG', 'KGZ', '417', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (120, 1604771561985900570, '老挝', 'Lao People\'s Democratic Republic (the)', 'the Lao People\'s Democratic Republic', 'LA', 'LAO', '418', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (121, 1604771561985900571, '拉脱维亚', 'Latvia', 'the Republic of Latvia', 'LV', 'LVA', '428', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (122, 1604771561985900572, '黎巴嫩', 'Lebanon', 'the Lebanese Republic', 'LB', 'LBN', '422', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (123, 1604771561985900573, '莱索托', 'Lesotho', 'the Kingdom of Lesotho', 'LS', 'LSO', '426', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (124, 1604771561985900574, '利比里亚', 'Liberia', 'the Republic of Liberia', 'LR', 'LBR', '430', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (125, 1604771561985900575, '利比亚', 'Libyan Arab Jamahiriya (the)', 'the Socialist People\'s Libyan Arab Jamahiriya', 'LY', 'LBY', '434', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (126, 1604771561985900576, '列支敦士登', 'Liechtenstein', 'the Principality of Liechtenstein', 'LI', 'LIE', '438', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (127, 1604771561985900577, '立陶宛', 'Lithuania', 'the Republic of Lithuania', 'LT', 'LTU', '440', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (128, 1604771561985900578, '卢森堡', 'Luxembourg', 'the Grand Duchy of Luxembourg', 'LU', 'LUX', '442', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (129, 1604771561985900579, '澳门', 'Macao', 'Macao Special Administrative Region of China', 'MO', 'MAC', '446', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (130, 1604771561985900580, '前南马其顿', 'Macedonia (the former Yugoslav Republic of)', 'the former Yugoslav Republic of Macedonia', 'MK', 'MKD', '807', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (131, 1604771561985900581, '马达加斯加', 'Madagascar', 'the Republic of Madagascar', 'MG', 'MDG', '450', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (132, 1604771561985900582, '马拉维', 'Malawi', 'the Republic of Malawi', 'MW', 'MWI', '454', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (133, 1604771561985900583, '马来西亚', 'Malaysia', '', 'MY', 'MYS', '458', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (134, 1604771561985900584, '马尔代夫', 'Maldives', 'the Republic of Maldives', 'MV', 'MDV', '462', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (135, 1604771561985900585, '马里', 'Mali', 'the Republic of Mali', 'ML', 'MLI', '466', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (136, 1604771561985900586, '马耳他', 'Malta', 'the Republic of Malta', 'MT', 'MLT', '470', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (137, 1604771561985900587, '马绍尔群岛', 'Marshall Islands (the)', 'the Republic of the Marshall Islands', 'MH', 'MHL', '584', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (138, 1604771561985900588, '马提尼克', 'Martinique', '', 'MQ', 'MTQ', '474', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (139, 1604771561985900589, '毛利塔尼亚', 'Mauritania', 'the Islamic Republic of Mauritania', 'MR', 'MRT', '478', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (140, 1604771561985900590, '毛里求斯', 'Mauritius', 'the Republic of Mauritius', 'MU', 'MUS', '480', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (141, 1604771561985900591, '马约特', 'Mayotte', '', 'YT', 'MYT', '175', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (142, 1604771561985900592, '墨西哥', 'Mexico', 'the United Mexican States', 'MX', 'MEX', '484', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (143, 1604771561985900593, '密克罗尼西亚联邦', 'Micronesia (the Federated States of)', 'the Federated States of Micronesia', 'FM', 'FSM', '583', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (144, 1604771561985900594, '摩尔多瓦', 'Moldova (the Republic of)', 'the Republic of Moldova', 'MD', 'MDA', '498', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (145, 1604771561985900595, '摩纳哥', 'Monaco', 'the Principality of Monaco', 'MC', 'MCO', '492', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (146, 1604771561985900596, '蒙古', 'Mongolia', '', 'MN', 'MNG', '496', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (147, 1604771561985900597, '黑山', 'Montenegro', 'he Republic of Montenegro', 'ME', 'MNE', '499', 'ISO 3166.1:2006新增', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (148, 1604771561985900598, '蒙特塞拉特', 'Montserrat', '', 'MS', 'MSR', '500', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (149, 1604771561985900599, '摩洛哥', 'Morocco', 'the Kingdom of Morocco', 'MA', 'MAR', '504', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (150, 1604771561985900600, '莫桑比克', 'Mozambique', 'the Republic of Mozambique', 'MZ', 'MOZ', '508', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (151, 1604771561985900601, '缅甸', 'Myanmar', 'the Union of Myanmar', 'MM', 'MMR', '104', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (152, 1604771561985900602, '纳米比亚', 'Namibia', 'the Republic of Namibia', 'NA', 'NAM', '516', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (153, 1604771561985900603, '瑙鲁', 'Nauru', 'the Republic of Nauru', 'NR', 'NRU', '520', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (154, 1604771561985900604, '尼泊尔', 'Nepal', 'Nepal', 'NP', 'NPL', '524', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (155, 1604771561985900605, '荷兰', 'Netherlands (the)', 'the Kingdom of the Netherlands', 'NL', 'NLD', '528', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (156, 1604771561985900606, '荷属安的列斯', 'Netherlands Antilles (the)', '', 'AN', 'ANT', '530', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (157, 1604771561985900607, '新喀里多尼亚', 'New Caledonia', '', 'NC', 'NCL', '540', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (158, 1604771561985900608, '新西兰', 'New Zealand', '', 'NZ', 'NZL', '554', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (159, 1604771561985900609, '尼加拉瓜', 'Nicaragua', 'the Republic of Nicaragua', 'NI', 'NIC', '558', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (160, 1604771561985900610, '尼日尔', 'Niger (the)', 'the Republic of the Niger', 'NE', 'NER', '562', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (161, 1604771561985900611, '尼日利亚', 'Nigeria', 'the Federal Republic of Nigeria', 'NG', 'NGA', '566', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (162, 1604771561985900612, '纽埃', 'Niue', 'the Republic of Niue', 'NU', 'NIU', '570', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (163, 1604771561985900613, '诺福克岛', 'Norfolk Island', '', 'NF', 'NFK', '574', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (164, 1604771561985900614, '北马里亚纳', 'Northern Mariana Islands (the)', 'the Commonwealth of the Northern Mariana Islands', 'MP', 'MNP', '580', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (165, 1604771561985900615, '挪威', 'Norway', 'the Kingdom of Norway', 'NO', 'NOR', '578', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (166, 1604771561990094848, '阿曼', 'Oman', 'the Sultanate of Oman', 'OM', 'OMN', '512', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (167, 1604771561990094849, '巴基斯坦', 'Pakistan', 'the Islamic Republic of Pakistan', 'PK', 'PAK', '586', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (168, 1604771561990094850, '帕劳', 'Palau', 'the Republic of Palau', 'PW', 'PLW', '585', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (169, 1604771561990094851, '巴勒斯坦', 'Palestinian Territory (the Occupied)', 'the Occupied Palestinian Territory', 'PS', 'PSE', '275', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (170, 1604771561990094852, '巴拿马', 'Panama', 'the Republic of Panama', 'PA', 'PAN', '591', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (171, 1604771561990094853, '巴布亚新几内亚', 'Papua New Guinea', '', 'PG', 'PNG', '598', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (172, 1604771561990094854, '巴拉圭', 'Paraguay', 'the Republic of Paraguay', 'PY', 'PRY', '600', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (173, 1604771561990094855, '秘鲁', 'Peru', 'the Republic of Peru', 'PE', 'PER', '604', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (174, 1604771561990094856, '菲律宾', 'Philippines (the)', 'the Republic of the Philippines', 'PH', 'PHL', '608', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (175, 1604771561990094857, '皮特凯恩', 'Pitcairn', '', 'PN', 'PCN', '612', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (176, 1604771561990094858, '波兰', 'Poland', 'the Republic of Poland', 'PL', 'POL', '616', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (177, 1604771561990094859, '葡萄牙', 'Portugal', 'the Portuguese Republic', 'PT', 'PRT', '620', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (178, 1604771561990094860, '波多黎各', 'Puerto Rico', '', 'PR', 'PRI', '630', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (179, 1604771561990094861, '卡塔尔', 'Qatar', 'the State of Qatar', 'QA', 'QAT', '634', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (180, 1604771561990094862, '留尼汪', 'Réunion', '', 'RE', 'REU', '638', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (181, 1604771561990094863, '罗马尼亚', 'Romania', '', 'RO', 'ROU', '642', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (182, 1604771561990094864, '俄罗斯联邦', 'Russian Federation (the)', 'the Russian Federation', 'RU', 'RUS', '643', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (183, 1604771561990094865, '卢旺达', 'Rwanda', 'the Republic of Rwanda', 'RW', 'RWA', '646', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (184, 1604771561994289152, '圣赫勒拿', 'Saint Helena', '', 'SH', 'SHN', '654', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (185, 1604771561994289153, '圣基茨和尼维斯', 'Saint Kitts and Nevis', '', 'KN', 'KNA', '659', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (186, 1604771561994289154, '圣卢西亚', 'Saint Lucia', '', 'LC', 'LCA', '662', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (187, 1604771561994289155, '圣皮埃尔和密克隆', 'Saint Pierre and Miquelon', '', 'PM', 'SPM', '666', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (188, 1604771561994289156, '圣文森特和格林纳丁斯', 'Saint Vincent and the Grenadines', '', 'VC', 'VCT', '670', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (189, 1604771561994289157, '萨摩亚', 'Samoa', 'the Independent State of Samoa', 'WS', 'WSM', '882', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (190, 1604771561994289158, '圣马力诺', 'San Marino', 'the Republic of San Marino', 'SM', 'SMR', '674', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (191, 1604771561994289159, '圣多美和普林西比', 'Sao Tome and Principe', 'the Democratic Republic of Sao Tome and Principe', 'ST', 'STP', '678', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (192, 1604771561994289160, '沙特阿拉伯', 'Saudi Arabia', 'the Kingdom of Saudi Arabia', 'SA', 'SAU', '682', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (193, 1604771561994289161, '塞内加尔', 'Senegal', 'the Republic of Senegal', 'SN', 'SEN', '686', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (194, 1604771561994289162, '塞尔维亚', 'Serbia', 'the Republic of Serbia', 'RS', 'SRB', '688', 'ISO 3166.1-2006新增', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (195, 1604771561994289163, '塞舌尔', 'Seychelles', 'the Republic of Seychelles', 'SC', 'SYC', '690', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (196, 1604771561994289164, '塞拉利昂', 'Sierra Leone', 'the Republic of Sierra Leone', 'SL', 'SLE', '694', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (197, 1604771561994289165, '新加坡', 'Singapore', 'the Republic of Singapore', 'SG', 'SGP', '702', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (198, 1604771561994289166, '斯洛伐克', 'Slovakia', 'the Slovak Republic', 'SK', 'SVK', '703', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (199, 1604771561994289167, '斯洛文尼亚', 'Slovenia', 'the Republic of Slovenia', 'SI', 'SVN', '705', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (200, 1604771561994289168, '所罗门群岛', 'Solomon Islands (the)', '', 'SB', 'SLB', '090', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (201, 1604771561994289169, '索马里', 'Somalia', 'the Somali Republic', 'SO', 'SOM', '706', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (202, 1604771561994289170, '南非', 'South Africa', 'the Republic of South Africa', 'ZA', 'ZAF', '710', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (203, 1604771561994289171, '南乔治亚岛和南桑德韦奇岛', 'South Georgia and the South Sandwich Islands', '', 'GS', 'SGS', '239', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (204, 1604771561994289172, '西班牙', 'Spain', 'the Kingdom of Spain', 'ES', 'ESP', '724', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (205, 1604771561994289173, '斯里兰卡', 'Sri Lanka', 'the Democratic Socialist Republic of Sri Lanka', 'LK', 'LKA', '144', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (206, 1604771561994289174, '苏丹', 'Sudan (the)', 'the Republic of the Sudan', 'SD', 'SDN', '736', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (207, 1604771561994289175, '苏里南', 'Suriname', 'the Republic of Suriname', 'SR', 'SUR', '740', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (208, 1604771561994289176, '斯瓦尔巴岛和扬马延岛', 'Svalbard and Jan Mayen', '', 'SJ', 'SJM', '744', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (209, 1604771561994289177, '斯威士兰', 'Swaziland', 'the Kingdom of Swaziland', 'SZ', 'SWZ', '748', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (210, 1604771561994289178, '瑞典', 'Sweden', 'the Kingdom of Sweden', 'SE', 'SWE', '752', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (211, 1604771561994289179, '瑞士', 'Switzerland', 'the Swiss Confederation', 'CH', 'CHE', '756', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (212, 1604771561994289180, '叙利亚', 'Syrian Arab Republic (the)', 'the Syrian Arab Republic', 'SY', 'SYR', '760', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (213, 1604771561994289181, '台湾', 'Taiwan (Province of China)', '', 'TW', 'TWN', '158', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (214, 1604771561994289182, '塔吉克斯坦', 'Tajikistan', 'the Republic of Tajikistan', 'TJ', 'TJK', '762', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (215, 1604771561994289183, '坦桑尼亚', 'Tanzania,United Republic of', 'the United Republic of Tanzania', 'TZ', 'TZA', '834', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (216, 1604771561994289184, '泰国', 'Thailand', 'the Kingdom of Thailand', 'TH', 'THA', '764', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (217, 1604771561994289185, '东帝汶', 'Timor-Leste', 'the Democratic Republic of Timor-Leste', 'TL', 'TLS', '626', 'ISO 3166.1-2006调整了英文名称和字母代码（原代码为TP\\TMP）', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (218, 1604771561994289186, '多哥', 'Togo', 'the Togolese Republic', 'TG', 'TGO', '768', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (219, 1604771561994289187, '托克劳', 'Tokelau', '', 'TK', 'TKL', '772', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (220, 1604771561994289188, '汤加', 'Tonga', 'the Kingdom of Tonga', 'TO', 'TON', '776', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (221, 1604771561994289189, '特立尼达和多巴哥', 'Trinidad and Tobago', 'the Republic of Trinidad and Tobago', 'TT', 'TTO', '780', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (222, 1604771561994289190, '突尼斯', 'Tunisia', 'the Republic of Tunisia', 'TN', 'TUN', '788', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (223, 1604771561994289191, '土耳其', 'Turkey', 'the Republic of Turkey', 'TR', 'TUR', '792', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (224, 1604771561994289192, '土库曼斯坦', 'Turkmenistan', '', 'TM', 'TKM', '795', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (225, 1604771561994289193, '特克斯和凯科斯群岛', 'Turks and Caicos Islands (the)', '', 'TC', 'TCA', '796', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (226, 1604771561994289194, '图瓦卢', 'Tuvalu', '', 'TV', 'TUV', '798', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (227, 1604771561994289195, '乌干达', 'Uganda', 'the Republic of Uganda', 'UG', 'UGA', '800', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (228, 1604771561994289196, '乌克兰', 'Ukraine', '', 'UA', 'UKR', '804', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (229, 1604771561994289197, '阿联酋', 'United Arab Emirates (the)', 'the United Arab Emirates', 'AE', 'ARE', '784', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (230, 1604771561994289198, '英国', 'United Kingdom (the)', 'the United Kingdom of Great Britain and Northern Ireland', 'GB', 'GBR', '826', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (231, 1604771561994289199, '美国', 'United States (the)', 'the United States of America', 'US', 'USA', '840', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (232, 1604771561994289200, '美国本土外小岛屿', 'United States Minor Outlying Islands (the)', '', 'UM', 'UMI', '581', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (233, 1604771561994289201, '乌拉圭', 'Uruguay', 'the Eastern Republic of Uruguay', 'UY', 'URY', '858', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (234, 1604771561994289202, '乌兹别克斯坦', 'Uzbekistan', 'the Republic of Uzbekistan', 'UZ', 'UZB', '860', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (235, 1604771561994289203, '瓦努阿图', 'Vanuatu', 'the Republic of Vanuatu', 'VU', 'VUT', '548', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (236, 1604771561994289204, '委内瑞拉', 'Venezuela', 'the Bolivarian Republic of Venezuela', 'VE', 'VEN', '862', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (237, 1604771561994289205, '越南', 'Viet Nam', 'the Socialist Republic of Viet Nam', 'VN', 'VNM', '704', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (238, 1604771561994289206, '英属维尔京群岛', 'Virgin Islands (British)', 'British Virgin Islands (the)', 'VG', 'VGB', '092', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (239, 1604771561994289207, '美属维尔京群岛', 'Virgin Islands (U.S.)', 'the Virgin Islands of the United States', 'VI', 'VIR', '850', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (240, 1604771561994289208, '瓦利斯和富图纳', 'Wallis and Futuna', 'Wallis and Futuna Islands', 'WF', 'WLF', '876', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (241, 1604771561994289209, '西撒哈拉', 'Western Sahara', '', 'EH', 'ESH', '732', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (242, 1604771561994289210, '也门', 'Yemen', 'the Republic of Yemen', 'YE', 'YEM', '887', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (243, 1604771561994289211, '赞比亚', 'Zambia', 'the Republic of Zambia', 'ZM', 'ZMB', '894', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
INSERT INTO `territory_code` VALUES (244, 1604771561994289212, '津巴布韦', 'Zimbabwe', 'the Republic of Zimbabwe', 'ZW', 'ZWE', '716', '', 0, '2022-12-19 17:32:10', '2022-12-19 17:32:32');
COMMIT;

-- ----------------------------
-- Table structure for user_score
-- ----------------------------
DROP TABLE IF EXISTS `user_score`;
CREATE TABLE `user_score` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` varchar(255) DEFAULT NULL,
  `score_result` int DEFAULT NULL,
  `zero_timestamp` bigint DEFAULT NULL,
  `create_time` bigint DEFAULT NULL,
  `update_time` bigint DEFAULT NULL,
  `score` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `row_unique` (`user_id`,`score_result`,`zero_timestamp`,`score`)
) ENGINE=InnoDB AUTO_INCREMENT=1072 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user_score
-- ----------------------------
BEGIN;
INSERT INTO `user_score` VALUES (1, '1000150', 1, 1629297601000, 1629297601000, 1629297601000, 1);
INSERT INTO `user_score` VALUES (4, '1000150', 45, 1629297601000, 1629297601000, 1649131381000, 100);
INSERT INTO `user_score` VALUES (6, '1000150', 1111, 1649130841000, 1649130841000, 1649130841000, 55);
INSERT INTO `user_score` VALUES (7, '1000150', 1111, 1649130853000, 1649130853000, 1649130853000, 55);
INSERT INTO `user_score` VALUES (8, '1000150', 1111, 1649131085000, 1649131085000, 1649131085000, 55);
INSERT INTO `user_score` VALUES (1001, '111', 222, 1649130841000, NULL, NULL, 20);
INSERT INTO `user_score` VALUES (1002, '111', 222, 1649130841000, NULL, NULL, 21);
INSERT INTO `user_score` VALUES (1009, '111', 222, 1649130841000, 1649130853000, 1649130853000, 29);
INSERT INTO `user_score` VALUES (1065, 'Jack', NULL, NULL, NULL, NULL, 18);
INSERT INTO `user_score` VALUES (1066, 'Jack', NULL, NULL, NULL, NULL, 18);
INSERT INTO `user_score` VALUES (1067, 'Jack', NULL, NULL, NULL, NULL, 18);
INSERT INTO `user_score` VALUES (1071, '22', 1111, 1649130841000, 1715414364000, 1715414364000, 55);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
