package honghua

/*
梭哈游戏用的是扑克牌共52张牌。在民间因为不容易出好牌，也有去掉234567的简易玩法（五张牌梭哈）。
牌型比较：同花顺>四条>满堂红>同花>顺子>三条>二对>单对>散牌。
数字比较：A>K>Q>J>10>9>8 >7>6>5>4>3>2
花式比较：黑桃>红桃>草花>方块

同花大顺  A心 K心 Q心 J心 10心
四条     4 4 4 4 9
满堂红    满堂红（Fullhouse，亦称“俘虏”、“骷髅”、“夫佬”、“葫芦”、“富尔豪斯”）：三张同一点数的牌，加一对其他点数的牌。
同花      五张同一花色的牌。
顺子      五张顺连的牌。
三条     4 4 4 9 10
两对     4 4 5 5 9
一对     4 4 2 3 6
无对     2 5 6 8 9


玩法:
1、常用术语：
（1）全梭：以最小玩家的金币数目为每个玩家梭哈时下注的最大数目，但是最大下注数目由房间确定。
（2）封顶：以最小玩家的金币数目的50%为每个玩家梭哈时下注的最大数目。但是最大下注数目依然为房间确定。最高封顶为 100万金币。
2、先发给各家一张底牌，底牌除本人外，要到决胜负时才可翻开。
3、从发第二张牌开始，每发一张牌，以牌面发展最佳者为优先，进行下注。
4、有人下注，想继续玩下去的人，要按“跟注”键，跟注后会下注到和上家相同的筹码，或可选择加注。根据房间的设定，可以在特定的时间选择“梭”，梭哈是加入桌面允许的最大下注。
5、各家如果觉得自己的牌况不妙，不想继续，可以按“放弃”键放弃下注，先前跟过的筹码，亦无法取回。
6、牌面最大的人可赢得桌面所有的筹码。当多家放弃，已经下的注不能收回，并且赢家的底牌不掀开。
7、纸牌种类：港式五张牌游戏用的是扑克牌，取各门花色的牌中的“8、9、10、J、Q、K、A”，共28张牌。


 */