package utils

import (
    "crypto/sha1"
    "fmt"
    "math"
    "time"
)

// hash a string
func PasswordHash(pwd string) string {
    hasher := sha1.New()
    hasher.Write([]byte(pwd))
    return fmt.Sprintf("%x", hasher.Sum(nil))
}

/** 微博时间格式化显示
 * @param timestamp，标准时间戳
 */
func SmcTimeSince(timeAt time.Time) string {
    now := time.Now()
    since := math.Abs(float64(now.Unix() - timeAt.Unix()))

    output := ""
    switch {
    case since < 60:
        output = "刚刚"
    case since < 60*60:
        output = fmt.Sprintf("%v分钟前", math.Floor(since/60))
    case since < 60*60*24:
        output = fmt.Sprintf("%v小时前", math.Floor(since/3600))
    case since < 60*60*24*2:
        output = fmt.Sprintf("昨天%v", timeAt.Format("15:04"))
    case since < 60*60*24*3:
        output = fmt.Sprintf("前天%v", timeAt.Format("15:04"))
    case timeAt.Format("2006") == now.Format("2006"):
        output = timeAt.Format("1月2日 15:04")
    default:
        output = timeAt.Format("2006年1月2日 15:04")
    }
    // if math.Floor(since/3600) > 0 {
    //     if timeAt.Format("2006-01-02") == now.Format("2006-01-02") {
    //         output = "今天 "
    //         output += timeAt.Format("15:04")
    //     } else {
    //         if timeAt.Format("2006") == now.Format("2006") {
    //             output = timeAt.Format("1月2日 15:04")
    //         } else {
    //             output = timeAt.Format("2006年1月2日 15:04")
    //         }
    //     }
    // } else {
    //     m := math.Floor(since / 60)
    //     if m > 0 {
    //         output = fmt.Sprintf("%v分钟前", m)
    //     } else {
    //         output = "刚刚"
    //     }
    // }
    return output
}
