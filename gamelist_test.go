package main

import (
	"encoding/xml"
	"fmt"
	"testing"
)

var raw = `
<gameList>
  <game>
    <path>xjqxz.zip</path>
    <name>仙剑奇侠传</name>
    <developer>仙剑奇侠传</developer>
    <desc>《仙剑奇侠传》是由大宇资讯所制作的一款国产单机中文角色扮演电脑游戏。本作是《仙剑奇侠传》系列的第一部作品，主题是“宿命”，于1995年7月发行。\n\n游戏讲述了平凡的客栈小伙计李逍遥因为机缘巧合结识了女娲族后裔赵灵儿。赵灵儿所居之仙灵岛意外遭人袭击，李逍遥遂担任起护送灵儿往苗疆寻母的重任。在旅程中，李逍遥先后结识了林月如、阿奴。在赵灵儿帮助苗人祈雨解旱、惩奸除恶的过程之中，赵灵儿的身世之谜以及她母亲的下落，也终究水落石出。此时，拜月教主的阴谋摆在众人的眼前，李逍遥、赵灵儿等人却不知这与拜月教主一战，将会铸成一段永恒的悲剧。\n\n该游戏曾获1995年CEM STAR“最佳角色扮演游戏奖"以及1995年KING TITLE"游戏类金袋奖"，并于2004年被改编成同名电视剧。</desc>
    <players>1</players>
  </game>
  <game>
    <path>jyqxz.zip</path>
    <name>金庸群侠传</name>
    <developer>金庸群侠传</developer>
    <desc>《金庸群侠传》是一款由智冠公司河洛工作室1996年推出的角色扮演类游戏，也是河洛工作室“群侠”系列三部曲的的第一部。第二部为为武林群侠传，第三部为三国群侠传。金庸群侠传于1996年发行，是由河洛工作室开发﹙后来改名为东方演算﹚、智冠科技发行的中文角色扮演游戏，在当时是一款相当受欢迎的游戏。\n\n此游戏把金庸先生“飞雪连天射白鹿，笑书神侠倚碧鸳”等14部小说中一些脍炙人口的片段，如六大门派围攻光明顶，冰火岛，英雄大会等融合在一起，改编而成，具有14个支线任务。玩家可自由选择游戏路线，其间基本上没有因果关系，只不过其中有一些任务相对比较难，因此需先完成别的任务才成。游戏的剧情是开放的，游戏地图也是开放的。在游戏中玩家可邀5位非主角控制的人物(NPC)，参加队伍，而且游戏中如果行善做好事多，则玩家的品德属性值高，这样一些人物才愿意与你为伍。反之，如果杀戮过重，最后过关时将遇到正派大侠们连手攻击，就较难回到现实世界里来了。</desc>
    <players>1</players>
  </game>
<gameList>
`

func TestDecode(t *testing.T) {
	var dat GameList
	xml.Unmarshal([]byte(raw), &dat)
	out, _ := xml.MarshalIndent(dat, "", " ")
	fmt.Println(string(out))
}
