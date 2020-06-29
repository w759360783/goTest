package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
√接口：视频，方法：制作，上映
√接口：黄暴产品，方法：刺激你的神经
√封装：封装电影类，属性：名字，公司，主演们，实现影视作品接口
√继承：爱情片、科幻片继承至电影，各自复写制作方法
√接口继承：东方艺术，继承影视产品、黄暴产品
√多态：影视作品细分：电视剧、网络剧
√类型断言：零点以前看随机电影、零点以后看东方艺术
*/

func main() {
	videos := GetVideos()
	hour := getRandomTime(23)
	fmt.Printf("当前时间为%d:00\n",hour)
	if hour > 9 && hour < 23 {
		//看随机电影
		for _, video := range videos {
			if isOrientalArt(video) {
				//东方艺术
				fmt.Println(video.GetName(),"不能看")
			} else {
				//非东方艺术
				fmt.Println("看",video.GetName())
			}
		}
	} else {
		//看东方艺术
		for _, video := range videos {
			if !isOrientalArt(video) {
				//非东方艺术
				fmt.Println(video.GetName(),"不能看")
			} else {
				//东方艺术
				fmt.Println("看",video.GetName())
			}
		}
	}
}

//鉴定视频是否为东方艺术

func isOrientalArt(video IVideo) bool {
	_, ok := video.(IOrientalArt)
	return ok
}

// 获取当前随机时间
func getRandomTime(s int) int {
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	return r.Intn(s)
}

//生成片库
func GetVideos() []IVideo {
	videos := make([]IVideo, 0)
	videos = append(
		videos, &AffectionalFilm{Film{Name: "泰坦尼克号", Company: "Fawkes",
			Star: []string{"小李子", "一个已经被人遗忘的中年妇女"}}})
	videos = append(
		videos, &ScienceFictionFilm{Film{Name: "星球大战", Company: "Lucas",
			Star: []string{"马克·哈米尔", "凯丽·费雪"}}})
	var ioa IOrientalArt
	laf := &LoveActionFilm{AffectionalFilm{Film{Name: "雨夜的撸嘛痴汉.avi", Company: "未知的公司",
		Star: []string{"小李子"}}}}
	ioa = laf
	videos = append(videos, ioa)
	return videos
}

//接口：视频，方法：制作、上映

type IVideo interface {
	//制作
	Make()
	//上映
	Release()
	//获取电影名
	GetName()string
}

//接口：黄暴产品，方法：刺激你的神经
type IYellowVideo interface {
	StimulatingNerve()
}

//接口继承：东方艺术，继承影视产品、黄暴产品
type IOrientalArt interface {
	IVideo
	IYellowVideo
}

//封装：封装电影类，属性：名字，公司，主演们，实现影视公司接口

type Film struct {
	Name    string
	Company string
	Star    []string
}

func (f *Film) Make() {
	fmt.Println("电影制作中......")
}

func (f *Film) Release() {
	fmt.Println("电影上映啦")
}

func (f *Film) GetName()string{
	return f.Name
}

//爱情片
type AffectionalFilm struct {
	Film
}

// 重写电影制作
func (af *AffectionalFilm) Make() {
	fmt.Println("爱情电影制作中")
}

// 科幻片
type ScienceFictionFilm struct {
	Film
}

// 重写电影制作
func (sff *ScienceFictionFilm) Make() {
	fmt.Println("科幻电影制作中")
}

//实现东方艺术实现
type LoveActionFilm struct {
	AffectionalFilm
}

func (laf *LoveActionFilm) StimulatingNerve() {
	fmt.Println("刺激你的神经")
}
