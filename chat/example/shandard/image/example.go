package main

import (
	"chat/pkg/log"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"math"
	"os"
)

// image实现了基本的2D图片库
// 基本接口叫作Image。图片的色彩定义在image/color包。
// Image接口可以通过调用如NewRGBA和NewPaletted函数等获得；也可以通过调用Decode函数解码包含GIF、JPEG或PNG格式图像数据的输入流获得
// 解码任何具体图像类型之前都必须注册对应类型的解码函数
// 注册过程一般是作为包初始化的副作用，放在包的init函数里。因此，要解码PNG图像，只需在程序的main包里嵌入  import _ "image/png"
func main() {
	// png图片生成
	examplePngEncode()
	// png图片解析
	examplePngDecode()

	exampleJpegEncode()
}

const (
	PngFilePath  = "example/testdata/rgb.png"
	JpegFilePath = "example/testdata/rgb.jpeg"
	GifFilePath  = "example/testdata/rgb.gif"
)

type Circle struct {
	X, Y, R float64
}

func (c *Circle) Brightness(x, y float64) uint8 {
	var dx, dy = c.X - x, c.Y - y
	d := math.Sqrt(dx*dx+dy*dy) / c.R
	if d > 1 {
		return 0
	} else {
		return 255
	}
}

func createRGBAImage() *image.RGBA {
	var w, h = 200, 240

	var hw, hh = float64(w / 2), float64(h / 2)
	r := 40.0

	p := 2 * math.Pi / 3
	cr := &Circle{hw - r*math.Sin(0), hh - r*math.Cos(0), 60}
	cg := &Circle{hw - r*math.Sin(p), hh - r*math.Cos(p), 60}
	cb := &Circle{hw - r*math.Sin(-p), hh-r*math.Cos(-p), 60}

	rect := image.Rect(0, 0, w, h)

	m := image.NewRGBA(rect)

	for x := 0; x < w; x++{
		for y := 0; y < h; y++ {
			c := color.RGBA{
				R: cr.Brightness(float64(x), float64(y)),
				G: cg.Brightness(float64(x), float64(y)),
				B: cb.Brightness(float64(x), float64(y)),
				A: 255,
			}
			m.Set(x, y, c)
		}
	}

	return m
}

func examplePngEncode()  {

	m := createRGBAImage()

	f, err := os.OpenFile(PngFilePath, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err = png.Encode(f, m); err != nil {
		log.Fatal(err)
	}
}

func examplePngDecode()  {

	f, err := os.Open(PngFilePath)
	if err != nil {
		log.Fatal(err)
	}

	img, err := png.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	levels := []string{" ", "░", "▒", "▓", "█"}

	for y := img.Bounds().Min.Y; y< img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			pointColor := img.At(x, y)

			c := color.GrayModel.Convert(pointColor).(color.Gray)
			level := c.Y / 51
			if level == 5 {
				level --
			}
			fmt.Print(levels[level])
		}
		fmt.Print("\n")
	}
}

func exampleJpegEncode()  {

	m := createRGBAImage()

	f, err := os.OpenFile(JpegFilePath, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	if err = jpeg.Encode(f, m, &jpeg.Options{Quality: 100}); err != nil {
		log.Fatal(err)
	}
}