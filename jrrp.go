package main

import (
	"fmt"
	"time"

	"github.com/catsworld/botmaid"
	"github.com/catsworld/random"
)

func init() {
	bm.AddTimer(botmaid.Timer{
		Do: func() {
			bm.Redis.Del("jrrp")
		},
		Start:     time.Date(2019, 10, 6, 0, 0, 0, 0, loc),
		Frequency: time.Hour * 24,
	})
	bm.AddCommand(&botmaid.Command{
		Do: func(u *botmaid.Update, b *botmaid.Bot) bool {
			rp := ""
			can := ""
			cannot := ""

			if bm.Redis.HExists("jrrp", fmt.Sprintf("%v_%v_rp", b.ID, u.User.ID)).Val() {
				rp = bm.Redis.HGet("jrrp", fmt.Sprintf("%v_%v_rp", b.ID, u.User.ID)).Val()
				can = bm.Redis.HGet("jrrp", fmt.Sprintf("%v_%v_can", b.ID, u.User.ID)).Val()
				cannot = bm.Redis.HGet("jrrp", fmt.Sprintf("%v_%v_cannot", b.ID, u.User.ID)).Val()
			} else {
				rp = fmt.Sprintf("%v", random.Int(1, 100))

				jrrpSkill := []string{
					"表演",
					"美术",
					"摄影",
					"伪造文书",
					"写作",
					"书法",
					"音乐",
					"厨艺",
					"表演",
					"理发",
					"木匠",
					"舞蹈",
					"莫里斯舞蹈",
					"歌剧演唱",
					"粉刷/油漆工",
					"制陶",
					"雕塑",
				}
				jrrpFighting := []string{
					"鞭子",
					"电锯",
					"斗殴",
					"斧",
					"剑",
					"绞索",
					"链枷",
					"矛",
				}
				jrrpShooting := []string{
					"步枪/霰弹枪",
					"冲锋枪",
					"弓",
					"火焰喷射器",
					"机枪",
					"手枪",
					"重武器",
				}
				jrrpScience := []string{
					"地质学",
					"化学",
					"生物学",
					"数学",
					"天文学",
					"物理学",
					"药学",
					"植物学",
					"动物学",
					"密码学",
					"工程学",
					"气象学",
					"司法科学",
				}
				jrrpSurvive := []string{
					"狩猎知识",
					"搭建住所",
					"危险意识",
				}
				jrrpCheckAndPK := []string{
					"力量检定",
					"力量对抗",
					"敏捷检定",
					"敏捷对抗",
					"智力检定",
					"灵感检定",
					"体质检定",
					"意志检定",
					"意志对抗",
					"教育检定",
					"幸运检定",
				}
				jrrpSpell := []string{
					"灵魂分配术",
					"耶德·艾塔德放逐术",
					"束缚术",
					"刀锋祝福术",
					"戈尔格罗斯形体扭曲术",
					"深渊之息",
					"黄金蜂蜜酒酿造术",
					"阿撒托斯请神术",
					"克图格亚请神术",
					"哈斯塔请神术",
					"伊塔库亚请神术",
					"尼约格萨请神术",
					"莎布-尼古拉斯请神术",
					"犹格-索托斯请神术",
					"阿撒托斯送神术",
					"克图格亚送神术",
					"哈斯塔送神术",
					"伊塔库亚送神术",
					"尼约格萨送神术",
					"莎布-尼古拉斯送神术",
					"犹格-索托斯送神术",
					"致盲术/治盲术",
					"透特之咏",
					"记忆模糊术",
					"尼约格萨紧握术",
					"外貌摄取术",
					"钻地魔虫联络术",
					"深潜者联络术",
					"古老者联络术",
					"飞水螅联络术",
					"无形之子联络术",
					"食尸鬼联络术",
					"诺弗-刻联络术",
					"廷达洛斯之猎犬联络术",
					"米-戈联络术",
					"人面鼠联络术",
					"潜沙怪联络术",
					"外神仆役联络术",
					"死灵联络术",
					"克苏鲁的星之眷属联络术",
					"伊斯人联络术",
					"昌格纳·方庚通神术",
					"克苏鲁通神术",
					"艾霍特通神术",
					"诺登斯通神术",
					"奈亚拉托提普通神术",
					"撒托古亚通神术",
					"伊戈罗纳克通神术",
					"纳克-提特障壁创建术",
					"拉莱耶造雾术",
					"僵尸制造术",
					"腐烂外皮之诅咒",
					"致死术",
					"支配术",
					"阿撒托斯的恐怖诅咒",
					"苏莱曼之尘",
					"旧印开光术",
					"书册附魔术",
					"刀具附魔术",
					"笛子附魔术",
					"祭刀附魔术",
					"哨子附魔术",
					"迷身术",
					"邪眼术",
					"犹格-索托斯之拳",
					"血肉防护术",
					"时空门",
					"时空门搜寻术",
					"时空箱创建术",
					"时光门",
					"时空门观察术",
					"绿腐术",
					"恐惧注入术",
					"血肉熔解术",
					"心理暗示术",
					"精神震爆术",
					"精神交换术",
					"精神转移术",
					"塔昆·阿提普之镜",
					"伊本-加兹之粉",
					"蒲林的埃及十字架",
					"修德·梅尔之赤印",
					"复活术",
					"枯萎术",
					"哈斯塔之歌",
					"拜亚基召唤术",
					"黑山羊幼崽召唤术",
					"空鬼召唤术",
					"炎之精召唤术",
					"恐怖猎手召唤术",
					"夜魇召唤术",
					"外神仆役召唤术",
					"星之精召唤术",
					"黑山羊幼崽束缚术",
					"空鬼束缚术",
					"炎之精束缚术",
					"恐怖猎手束缚术",
					"夜魇束缚术",
					"外神仆役束缚术",
					"星之精束缚术",
					"维瑞之印",
					"守卫术",
					"忘却之波",
					"肢体凋萎术",
					"真言术",
					"折磨术",
				}
				jrrpTRPG := []string{
					"守序善良",
					"中立善良",
					"混乱善良",
					"守序中立",
					"绝对中立",
					"混乱中立",
					"守序邪恶",
					"中立邪恶",
					"混乱邪恶",
				}

				jrrpEvent := []string{
					"会计",
					"人类学",
					"估价",
					"考古学",
					"技艺：" + random.String(jrrpSkill),
					"魅惑",
					"攀爬",
					"计算机使用",
					"信用评级",
					"克苏鲁神话",
					"乔装",
					"闪避",
					"汽车驾驶",
					"电气维修",
					"电子学",
					"话术",
					"格斗：" + random.String(jrrpFighting),
					"射击：" + random.String(jrrpShooting),
					"急救",
					"历史",
					"恐吓",
					"跳跃",
					"语言：外语",
					"母语",
					"法律",
					"图书馆使用",
					"聆听",
					"锁匠",
					"机械维修",
					"医学",
					"博物学",
					"领航",
					"神秘学",
					"操作重型机械",
					"说服",
					"驾驶：飞行器",
					"驾驶：船",
					"精神分析",
					"心理学",
					"骑术",
					"科学：" + random.String(jrrpScience),
					"妙手",
					"侦察",
					"潜行",
					"生存：" + random.String(jrrpSurvive),
					"游泳",
					"投掷",
					"追踪",
					"潜水",
					"爆破",
					"催眠",
					"读唇",
					"炮术",
					"驯兽",
					random.String(jrrpCheckAndPK),
					"阅读神话典籍",
					"战斗",
					"追逐",
					"法术：" + random.String(jrrpSpell),
					"孤注一掷",
					random.String(jrrpTRPG) + "跑团",
					"SAN check",
					"调戏小阿比",
				}

				can = random.String(jrrpEvent)
				cannot = random.String(jrrpEvent)
				for can == cannot {
					cannot = random.String(jrrpEvent)
				}

				bm.Redis.HSet("jrrp", fmt.Sprintf("%v_%v_rp", b.ID, u.User.ID), rp)
				bm.Redis.HSet("jrrp", fmt.Sprintf("%v_%v_can", b.ID, u.User.ID), can)
				bm.Redis.HSet("jrrp", fmt.Sprintf("%v_%v_cannot", b.ID, u.User.ID), cannot)
			}

			b.Reply(u, fmt.Sprintf((botmaid.WordSlice{
				botmaid.Word{
					Word:   "%v今日的人品值是——%v，宜%v，忌%v",
					Weight: 49,
				},
				botmaid.Word{
					Word:   "%v今天有%v点人品，宜%v，忌%v，可不要把人品值用光了哦XD",
					Weight: 1,
				},
			}.Random()), u.User.NickName, rp, can, cannot))
			return true
		},
		Menu:       "jrrp",
		MenuText:   "今日人品",
		Names:      []string{"jrrp"},
		ArgsMinLen: 1,
		ArgsMaxLen: 1,
		Help:       " - 占卜今日人品",
	})
}
