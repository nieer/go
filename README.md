根据传入日期获取相关日期包括今天、昨天、本周始末、本月始末、本季度始末和各个维度同比日期等
<br />
例如：<br />
timeFormat.Init("2015-07-01")<br />
timeFormat.Today() // 2015-07-01<br />
timeFormat.Yesterday() // 2015-06-30<br />
timeFormat.ThisWeekStart() // 2015-06-29<br />
timeFormat.ThisWeekEnd() // 2015-07-05<br />
...<br />
