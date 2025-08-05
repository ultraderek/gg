package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"

	"github.com/golang/freetype/truetype"
	"github.com/ultraderek/gg"
	"golang.org/x/image/font/gofont/goregular"
)

func main() {
	wrap("out/wrap.png")
	unicode("out/unicode.png")
	tiling("out/tiling.png")
	texting("out/text.png")
	stars("out/stars.png")
	star("out/star.png")
	spiral("out/spiral.png")
	sine("out/sine.png")
	scatter("out/scatter.png")
	rotadedtext("out/rotadedtext.png")
	rotatedimage("out/rotatedimage.png")
	quadratic("out/quadratic.png")
	patternfill("out/patternfill.png")
	openfill("out/openfill.png")
	meme("out/meme.png")
	mask("out/mask.png")
	lorem("out/lorem.png")
	linewidth("out/linewidth.png")
	linesEX("out/linesEX.png")
	invertmask("out/invertmask.png")
	gradienttext("out/gradienttext.png")
	gradientradial("out/gradientradial.png")
	gradientlinear("out/gradientlinear.png")
	gradientconic("out/gradientconic.png")
	gofont("out/gofont.png")
	ellipse("out/ellipse.png")
	cubic("out/cubic.png")
	crisp("out/crisp.png")
	concat("out/concat.png")
	clip("out/clip.png")
	circle("out/circle.png")
	beziers("out/beziers.png")
}
func beziers(filename string) {
	random := func() float64 {
		return rand.Float64()*2 - 1
	}

	point := func() (x, y float64) {
		return random(), random()
	}

	drawCurve := func(dc *gg.Context) {
		dc.SetRGBA(0, 0, 0, 0.1)
		dc.FillPreserve()
		dc.SetRGB(0, 0, 0)
		dc.SetLineWidth(12)
		dc.Stroke()
	}

	drawPoints := func(dc *gg.Context) {
		dc.SetRGBA(1, 0, 0, 0.5)
		dc.SetLineWidth(2)
		dc.Stroke()
	}

	randomQuadratic := func(dc *gg.Context) {
		x0, y0 := point()
		x1, y1 := point()
		x2, y2 := point()
		dc.MoveTo(x0, y0)
		dc.QuadraticTo(x1, y1, x2, y2)
		drawCurve(dc)
		dc.MoveTo(x0, y0)
		dc.LineTo(x1, y1)
		dc.LineTo(x2, y2)
		drawPoints(dc)
	}

	randomCubic := func(dc *gg.Context) {
		x0, y0 := point()
		x1, y1 := point()
		x2, y2 := point()
		x3, y3 := point()
		dc.MoveTo(x0, y0)
		dc.CubicTo(x1, y1, x2, y2, x3, y3)
		drawCurve(dc)
		dc.MoveTo(x0, y0)
		dc.LineTo(x1, y1)
		dc.LineTo(x2, y2)
		dc.LineTo(x3, y3)
		drawPoints(dc)
	}
	const (
		S = 256
		W = 8
		H = 8
	)
	dc := gg.NewContext(S*W, S*H)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	for j := 0; j < H; j++ {
		for i := 0; i < W; i++ {
			x := float64(i)*S + S/2
			y := float64(j)*S + S/2
			dc.Push()
			dc.Translate(x, y)
			dc.Scale(S/2, S/2)
			if j%2 == 0 {
				randomCubic(dc)
			} else {
				randomQuadratic(dc)
			}
			dc.Pop()
		}
	}
	dc.SavePNG(filename)
}
func circle(filename string) {
	dc := gg.NewContext(1000, 1000)
	dc.DrawCircle(500, 500, 400)
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	dc.SavePNG(filename)
}
func clip(filename string) {
	dc := gg.NewContext(1000, 1000)
	dc.DrawCircle(350, 500, 300)
	dc.Clip()
	dc.DrawCircle(650, 500, 300)
	dc.Clip()
	dc.DrawRectangle(0, 0, 1000, 1000)
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	dc.SavePNG(filename)
}
func concat(filename string) {
	im1, err := gg.LoadPNG("baboon.png")
	if err != nil {
		panic(err)
	}

	im2, err := gg.LoadPNG("gopher.png")
	if err != nil {
		panic(err)
	}

	s1 := im1.Bounds().Size()
	s2 := im2.Bounds().Size()

	width := int(math.Max(float64(s1.X), float64(s2.X)))
	height := s1.Y + s2.Y

	dc := gg.NewContext(width, height)
	dc.DrawImage(im1, 0, 0)
	dc.DrawImage(im2, 0, s1.Y)
	dc.SavePNG(filename)
}
func crisp(filename string) {
	const W = 1000
	const H = 1000
	const Minor = 10
	const Major = 100

	dc := gg.NewContext(W, H)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// minor grid
	for x := Minor; x < W; x += Minor {
		fx := float64(x) + 0.5
		dc.DrawLine(fx, 0, fx, H)
	}
	for y := Minor; y < H; y += Minor {
		fy := float64(y) + 0.5
		dc.DrawLine(0, fy, W, fy)
	}
	dc.SetLineWidth(1)
	dc.SetRGBA(0, 0, 0, 0.25)
	dc.Stroke()

	// major grid
	for x := Major; x < W; x += Major {
		fx := float64(x) + 0.5
		dc.DrawLine(fx, 0, fx, H)
	}
	for y := Major; y < H; y += Major {
		fy := float64(y) + 0.5
		dc.DrawLine(0, fy, W, fy)
	}
	dc.SetLineWidth(1)
	dc.SetRGBA(0, 0, 0, 0.5)
	dc.Stroke()

	dc.SavePNG(filename)
}
func cubic(filename string) {
	const S = 1000
	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.Translate(S/2, S/2)
	dc.Scale(40, 40)

	var x0, y0, x1, y1, x2, y2, x3, y3 float64
	x0, y0 = -10, 0
	x1, y1 = -8, -8
	x2, y2 = 8, 8
	x3, y3 = 10, 0

	dc.MoveTo(x0, y0)
	dc.CubicTo(x1, y1, x2, y2, x3, y3)
	dc.SetRGBA(0, 0, 0, 0.2)
	dc.SetLineWidth(8)
	dc.FillPreserve()
	dc.SetRGB(0, 0, 0)
	dc.SetDash(16, 24)
	dc.Stroke()

	dc.MoveTo(x0, y0)
	dc.LineTo(x1, y1)
	dc.LineTo(x2, y2)
	dc.LineTo(x3, y3)
	dc.SetRGBA(1, 0, 0, 0.4)
	dc.SetLineWidth(2)
	dc.SetDash(4, 8, 1, 8)
	dc.Stroke()

	dc.SavePNG(filename)
}
func ellipse(filename string) {
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetRGBA(0, 0, 0, 0.1)
	for i := 0; i < 360; i += 15 {
		dc.Push()
		dc.RotateAbout(gg.Radians(float64(i)), S/2, S/2)
		dc.DrawEllipse(S/2, S/2, S*7/16, S/8)
		dc.Fill()
		dc.Pop()
	}
	if im, err := gg.LoadImage("gopher.png"); err == nil {
		dc.DrawImageAnchored(im, S/2, S/2, 0.5, 0.5)
	}
	dc.SavePNG(filename)
}
func gofont(filename string) {
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}

	face := truetype.NewFace(font, &truetype.Options{Size: 48})

	dc := gg.NewContext(1024, 1024)
	dc.SetFontFace(face)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.DrawStringAnchored("Hello, world!", 512, 512, 0.5, 0.5)
	dc.SavePNG(filename)
}
func gradientconic(filename string) {
	dc := gg.NewContext(400, 400)

	grad1 := gg.NewConicGradient(200, 200, 0)
	grad1.AddColorStop(0.0, color.Black)
	grad1.AddColorStop(0.5, color.RGBA{255, 215, 0, 255})
	grad1.AddColorStop(1.0, color.RGBA{255, 0, 0, 255})

	grad2 := gg.NewConicGradient(200, 200, 90)
	grad2.AddColorStop(0.00, color.RGBA{255, 0, 0, 255})
	grad2.AddColorStop(0.16, color.RGBA{255, 255, 0, 255})
	grad2.AddColorStop(0.33, color.RGBA{0, 255, 0, 255})
	grad2.AddColorStop(0.50, color.RGBA{0, 255, 255, 255})
	grad2.AddColorStop(0.66, color.RGBA{0, 0, 255, 255})
	grad2.AddColorStop(0.83, color.RGBA{255, 0, 255, 255})
	grad2.AddColorStop(1.00, color.RGBA{255, 0, 0, 255})

	dc.SetStrokeStyle(grad1)
	dc.SetLineWidth(20)
	dc.DrawCircle(200, 200, 180)
	dc.Stroke()

	dc.SetFillStyle(grad2)
	dc.DrawCircle(200, 200, 150)
	dc.Fill()

	dc.SavePNG(filename)
}
func gradientlinear(filename string) {
	dc := gg.NewContext(500, 400)

	grad := gg.NewLinearGradient(20, 320, 400, 20)
	grad.AddColorStop(0, color.RGBA{0, 255, 0, 255})
	grad.AddColorStop(1, color.RGBA{0, 0, 255, 255})
	grad.AddColorStop(0.5, color.RGBA{255, 0, 0, 255})

	dc.SetColor(color.White)
	dc.DrawRectangle(20, 20, 400-20, 300)
	dc.Stroke()

	dc.SetStrokeStyle(grad)
	dc.SetLineWidth(4)
	dc.MoveTo(10, 10)
	dc.LineTo(410, 10)
	dc.LineTo(410, 100)
	dc.LineTo(10, 100)
	dc.ClosePath()
	dc.Stroke()

	dc.SetFillStyle(grad)
	dc.MoveTo(10, 120)
	dc.LineTo(410, 120)
	dc.LineTo(410, 300)
	dc.LineTo(10, 300)
	dc.ClosePath()
	dc.Fill()

	dc.SavePNG(filename)
}
func gradientradial(filename string) {
	dc := gg.NewContext(400, 200)

	grad := gg.NewRadialGradient(100, 100, 10, 100, 120, 80)
	grad.AddColorStop(0, color.RGBA{0, 255, 0, 255})
	grad.AddColorStop(1, color.RGBA{0, 0, 255, 255})

	dc.SetFillStyle(grad)
	dc.DrawRectangle(0, 0, 200, 200)
	dc.Fill()

	dc.SetColor(color.White)
	dc.DrawCircle(100, 100, 10)
	dc.Stroke()
	dc.DrawCircle(100, 120, 80)
	dc.Stroke()

	dc.SavePNG(filename)
}
func gradienttext(filename string) {
	const (
		W = 1024
		H = 512
	)
	dc := gg.NewContext(W, H)

	// draw text
	dc.SetRGB(0, 0, 0)
	dc.LoadFontFace("/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf", 128)
	dc.DrawStringAnchored("Gradient Text", W/2, H/2, 0.5, 0.5)

	// get the context as an alpha mask
	mask := dc.AsMask()

	// clear the context
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// set a gradient
	g := gg.NewLinearGradient(0, 0, W, H)
	g.AddColorStop(0, color.RGBA{255, 0, 0, 255})
	g.AddColorStop(1, color.RGBA{0, 0, 255, 255})
	dc.SetFillStyle(g)

	// using the mask, fill the context with the gradient
	dc.SetMask(mask)
	dc.DrawRectangle(0, 0, W, H)
	dc.Fill()

	dc.SavePNG(filename)
}
func invertmask(filename string) {
	dc := gg.NewContext(1024, 1024)
	dc.DrawCircle(512, 512, 384)
	dc.Clip()
	dc.InvertMask()
	dc.DrawRectangle(0, 0, 1024, 1024)
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	dc.SavePNG(filename)
}
func linesEX(filename string) {
	const W = 1024
	const H = 1024
	dc := gg.NewContext(W, H)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	for i := 0; i < 1000; i++ {
		x1 := rand.Float64() * W
		y1 := rand.Float64() * H
		x2 := rand.Float64() * W
		y2 := rand.Float64() * H
		r := rand.Float64()
		g := rand.Float64()
		b := rand.Float64()
		a := rand.Float64()*0.5 + 0.5
		w := rand.Float64()*4 + 1
		dc.SetRGBA(r, g, b, a)
		dc.SetLineWidth(w)
		dc.DrawLine(x1, y1, x2, y2)
		dc.Stroke()
	}
	dc.SavePNG(filename)
}
func linewidth(filename string) {
	dc := gg.NewContext(1000, 1000)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	w := 0.1
	for i := 100; i <= 900; i += 20 {
		x := float64(i)
		dc.DrawLine(x+50, 0, x-50, 1000)
		dc.SetLineWidth(w)
		dc.Stroke()
		w += 0.1
	}
	dc.SavePNG(filename)
}
func lorem(filename string) {
	var lines = []string{
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod",
		"tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,",
		"quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo",
		"consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse",
		"cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat",
		"non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
	}
	const W = 800
	const H = 400
	dc := gg.NewContext(W, H)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	// dc.LoadFontFace("/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf", 18)
	const h = 24
	for i, line := range lines {
		y := H/2 - h*len(lines)/2 + i*h
		dc.DrawStringAnchored(line, 400, float64(y), 0.5, 0.5)
	}
	dc.SavePNG(filename)
}
func mask(filename string) {
	im, err := gg.LoadImage("baboon.png")
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContext(512, 512)
	dc.DrawRoundedRectangle(0, 0, 512, 512, 64)
	dc.Clip()
	dc.DrawImage(im, 0, 0)
	dc.SavePNG(filename)
}
func meme(filename string) {
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	if err := dc.LoadFontFace("/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf", 96); err != nil {
		panic(err)
	}
	dc.SetRGB(0, 0, 0)
	s := "ONE DOES NOT SIMPLY"
	n := 6 // "stroke" size
	for dy := -n; dy <= n; dy++ {
		for dx := -n; dx <= n; dx++ {
			if dx*dx+dy*dy >= n*n {
				// give it rounded corners
				continue
			}
			x := S/2 + float64(dx)
			y := S/2 + float64(dy)
			dc.DrawStringAnchored(s, x, y, 0.5, 0.5)
		}
	}
	dc.SetRGB(1, 1, 1)
	dc.DrawStringAnchored(s, S/2, S/2, 0.5, 0.5)
	dc.SavePNG(filename)
}
func openfill(filename string) {
	dc := gg.NewContext(1000, 1000)
	for j := 0; j < 10; j++ {
		for i := 0; i < 10; i++ {
			x := float64(i)*100 + 50
			y := float64(j)*100 + 50
			a1 := rand.Float64() * 2 * math.Pi
			a2 := a1 + rand.Float64()*math.Pi + math.Pi/2
			dc.DrawArc(x, y, 40, a1, a2)
			// dc.ClosePath()
		}
	}
	dc.SetRGB(0, 0, 0)
	dc.FillPreserve()
	dc.SetRGB(1, 1, 1)
	dc.SetLineWidth(8)
	dc.StrokePreserve()
	dc.SetRGB(1, 0, 0)
	dc.SetLineWidth(4)
	dc.StrokePreserve()
	dc.SavePNG(filename)
}
func patternfill(filename string) {
	im, err := gg.LoadPNG("baboon.png")
	if err != nil {
		panic(err)
	}
	pattern := gg.NewSurfacePattern(im, gg.RepeatBoth)
	dc := gg.NewContext(600, 600)
	dc.MoveTo(20, 20)
	dc.LineTo(590, 20)
	dc.LineTo(590, 590)
	dc.LineTo(20, 590)
	dc.ClosePath()
	dc.SetFillStyle(pattern)
	dc.Fill()
	dc.SavePNG(filename)
}
func quadratic(filename string) {
	const S = 1000
	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.Translate(S/2, S/2)
	dc.Scale(40, 40)

	var x0, y0, x1, y1, x2, y2, x3, y3, x4, y4 float64
	x0, y0 = -10, 0
	x1, y1 = -5, -10
	x2, y2 = 0, 0
	x3, y3 = 5, 10
	x4, y4 = 10, 0

	dc.MoveTo(x0, y0)
	dc.LineTo(x1, y1)
	dc.LineTo(x2, y2)
	dc.LineTo(x3, y3)
	dc.LineTo(x4, y4)
	dc.SetHexColor("FF2D00")
	dc.SetLineWidth(8)
	dc.Stroke()

	dc.MoveTo(x0, y0)
	dc.QuadraticTo(x1, y1, x2, y2)
	dc.QuadraticTo(x3, y3, x4, y4)
	dc.SetHexColor("3E606F")
	dc.SetLineWidth(16)
	dc.FillPreserve()
	dc.SetRGB(0, 0, 0)
	dc.Stroke()

	dc.DrawCircle(x0, y0, 0.5)
	dc.DrawCircle(x1, y1, 0.5)
	dc.DrawCircle(x2, y2, 0.5)
	dc.DrawCircle(x3, y3, 0.5)
	dc.DrawCircle(x4, y4, 0.5)
	dc.SetRGB(1, 1, 1)
	dc.FillPreserve()
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(4)
	dc.Stroke()

	dc.LoadFontFace("/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf", 200)
	dc.DrawStringAnchored("g", -5, 5, 0.5, 0.5)
	dc.DrawStringAnchored("G", 5, -5, 0.5, 0.5)

	dc.SavePNG(filename)
}
func rotatedimage(filename string) {
	const W = 400
	const H = 500
	im, err := gg.LoadPNG("gopher.png")
	if err != nil {
		panic(err)
	}
	iw, ih := im.Bounds().Dx(), im.Bounds().Dy()
	dc := gg.NewContext(W, H)
	// draw outline
	dc.SetHexColor("#ff0000")
	dc.SetLineWidth(1)
	dc.DrawRectangle(0, 0, float64(W), float64(H))
	dc.Stroke()
	// draw full image
	dc.SetHexColor("#0000ff")
	dc.SetLineWidth(2)
	dc.DrawRectangle(100, 210, float64(iw), float64(ih))
	dc.Stroke()
	dc.DrawImage(im, 100, 210)
	// draw image with current matrix applied
	dc.SetHexColor("#0000ff")
	dc.SetLineWidth(2)
	dc.Rotate(gg.Radians(10))
	dc.DrawRectangle(100, 0, float64(iw), float64(ih)/2+20.0)
	dc.StrokePreserve()
	dc.Clip()
	dc.DrawImageAnchored(im, 100, 0, 0.0, 0.0)
	dc.SavePNG(filename)
}
func rotadedtext(filename string) {
	const S = 400
	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic("")
	}
	face := truetype.NewFace(font, &truetype.Options{
		Size: 40,
	})
	dc.SetFontFace(face)
	text := "Hello, world!"
	w, h := dc.MeasureString(text)
	dc.Rotate(gg.Radians(10))
	dc.DrawRectangle(100, 180, w, h)
	dc.Stroke()
	dc.DrawStringAnchored(text, 100, 180, 0.0, 0.0)
	dc.SavePNG(filename)
}
func scatter(filename string) {
	CreatePoints := func(n int) []gg.Point {
		points := make([]gg.Point, n)
		for i := 0; i < n; i++ {
			x := 0.5 + rand.NormFloat64()*0.1
			y := x + rand.NormFloat64()*0.1
			points[i] = gg.Point{x, y}
		}
		return points
	}
	const S = 1024
	const P = 64
	dc := gg.NewContext(S, S)
	dc.InvertY()
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	points := CreatePoints(1000)
	dc.Translate(P, P)
	dc.Scale(S-P*2, S-P*2)
	// draw minor grid
	for i := 1; i <= 10; i++ {
		x := float64(i) / 10
		dc.MoveTo(x, 0)
		dc.LineTo(x, 1)
		dc.MoveTo(0, x)
		dc.LineTo(1, x)
	}
	dc.SetRGBA(0, 0, 0, 0.25)
	dc.SetLineWidth(1)
	dc.Stroke()
	// draw axes
	dc.MoveTo(0, 0)
	dc.LineTo(1, 0)
	dc.MoveTo(0, 0)
	dc.LineTo(0, 1)
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(4)
	dc.Stroke()
	// draw points
	dc.SetRGBA(0, 0, 1, 0.5)
	for _, p := range points {
		dc.DrawCircle(p.X, p.Y, 3.0/S)
		dc.Fill()
	}
	// draw text
	dc.Identity()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf", 24); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored("Chart Title", S/2, P/2, 0.5, 0.5)
	if err := dc.LoadFontFace("/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf", 18); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored("X Axis Title", S/2, S-P/2, 0.5, 0.5)
	dc.SavePNG(filename)
}
func sine(filename string) {
	const W = 1200
	const H = 60
	dc := gg.NewContext(W, H)
	// dc.SetHexColor("#FFFFFF")
	// dc.Clear()
	dc.ScaleAbout(0.95, 0.75, W/2, H/2)
	for i := 0; i < W; i++ {
		a := float64(i) * 2 * math.Pi / W * 8
		x := float64(i)
		y := (math.Sin(a) + 1) / 2 * H
		dc.LineTo(x, y)
	}
	dc.ClosePath()
	dc.SetHexColor("#3E606F")
	dc.FillPreserve()
	dc.SetHexColor("#19344180")
	dc.SetLineWidth(8)
	dc.Stroke()
	dc.SavePNG(filename)
}
func spiral(filename string) {
	const S = 1024
	const N = 2048
	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	for i := 0; i <= N; i++ {
		t := float64(i) / N
		d := t*S*0.4 + 10
		a := t * math.Pi * 2 * 20
		x := S/2 + math.Cos(a)*d
		y := S/2 + math.Sin(a)*d
		r := t * 8
		dc.DrawCircle(x, y, r)
	}
	dc.Fill()
	dc.SavePNG(filename)
}
func star(filename string) {
	type Point struct {
		X, Y float64
	}

	Polygon := func(n int, x, y, r float64) []Point {
		result := make([]Point, n)
		for i := 0; i < n; i++ {
			a := float64(i)*2*math.Pi/float64(n) - math.Pi/2
			result[i] = Point{x + r*math.Cos(a), y + r*math.Sin(a)}
		}
		return result
	}
	n := 5
	points := Polygon(n, 512, 512, 400)
	dc := gg.NewContext(1024, 1024)
	dc.SetHexColor("fff")
	dc.Clear()
	for i := 0; i < n+1; i++ {
		index := (i * 2) % n
		p := points[index]
		dc.LineTo(p.X, p.Y)
	}
	dc.SetRGBA(0, 0.5, 0, 1)
	dc.SetFillRule(gg.FillRuleEvenOdd)
	dc.FillPreserve()
	dc.SetRGBA(0, 1, 0, 0.5)
	dc.SetLineWidth(16)
	dc.Stroke()
	dc.SavePNG(filename)
}
func stars(filename string) {
	type Point struct {
		X, Y float64
	}

	Polygon := func(n int) []Point {
		result := make([]Point, n)
		for i := 0; i < n; i++ {
			a := float64(i)*2*math.Pi/float64(n) - math.Pi/2
			result[i] = Point{math.Cos(a), math.Sin(a)}
		}
		return result
	}
	const W = 1200
	const H = 120
	const S = 100
	dc := gg.NewContext(W, H)
	dc.SetHexColor("#FFFFFF")
	dc.Clear()
	n := 5
	points := Polygon(n)
	for x := S / 2; x < W; x += S {
		dc.Push()
		s := rand.Float64()*S/4 + S/4
		dc.Translate(float64(x), H/2)
		dc.Rotate(rand.Float64() * 2 * math.Pi)
		dc.Scale(s, s)
		for i := 0; i < n+1; i++ {
			index := (i * 2) % n
			p := points[index]
			dc.LineTo(p.X, p.Y)
		}
		dc.SetLineWidth(10)
		dc.SetHexColor("#FFCC00")
		dc.StrokePreserve()
		dc.SetHexColor("#FFE43A")
		dc.Fill()
		dc.Pop()
	}
	dc.SavePNG(filename)
}
func texting(filename string) {
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf", 96); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored("Hello, world!", S/2, S/2, 0.5, 0.5)
	dc.SavePNG(filename)
}
func tiling(filename string) {
	const NX = 4
	const NY = 3
	im, err := gg.LoadPNG("gopher.png")
	if err != nil {
		panic(err)
	}
	w := im.Bounds().Size().X
	h := im.Bounds().Size().Y
	dc := gg.NewContext(w*NX, h*NY)
	for y := 0; y < NY; y++ {
		for x := 0; x < NX; x++ {
			dc.DrawImage(im, x*w, y*h)
		}
	}
	dc.SavePNG(filename)
}
func unicode(filename string) {
	const S = 4096 * 2
	const T = 16 * 2
	const F = 28
	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf", F); err != nil {
		panic(err)
	}
	for r := 0; r < 256; r++ {
		for c := 0; c < 256; c++ {
			i := r*256 + c
			x := float64(c*T) + T/2
			y := float64(r*T) + T/2
			dc.DrawStringAnchored(string(rune(i)), x, y, 0.5, 0.5)
		}
	}
	dc.SavePNG(filename)
}
func wrap(filename string) {
	const TEXT = "Call me Ishmael. Some years ago—never mind how long precisely—having little or no money in my purse, and nothing particular to interest me on shore, I thought I would sail about a little and see the watery part of the world. It is a way I have of driving off the spleen and regulating the circulation. Whenever I find myself growing grim about the mouth; whenever it is a damp, drizzly November in my soul; whenever I find myself involuntarily pausing before coffin warehouses, and bringing up the rear of every funeral I meet; and especially whenever my hypos get such an upper hand of me, that it requires a strong moral principle to prevent me from deliberately stepping into the street, and methodically knocking people's hats off—then, I account it high time to get to sea as soon as I can. This is my substitute for pistol and ball. With a philosophical flourish Cato throws himself upon his sword; I quietly take to the ship. There is nothing surprising in this. If they but knew it, almost all men in their degree, some time or other, cherish very nearly the same feelings towards the ocean with me."
	const W = 1024
	const H = 1024
	const P = 16
	dc := gg.NewContext(W, H)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.DrawLine(W/2, 0, W/2, H)
	dc.DrawLine(0, H/2, W, H/2)
	dc.DrawRectangle(P, P, W-P-P, H-P-P)
	dc.SetRGBA(0, 0, 1, 0.25)
	dc.SetLineWidth(3)
	dc.Stroke()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf", 18); err != nil {
		panic(err)
	}
	dc.DrawStringWrapped("UPPER LEFT", P, P, 0, 0, 0, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped("UPPER RIGHT", W-P, P, 1, 0, 0, 1.5, gg.AlignRight)
	dc.DrawStringWrapped("BOTTOM LEFT", P, H-P, 0, 1, 0, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped("BOTTOM RIGHT", W-P, H-P, 1, 1, 0, 1.5, gg.AlignRight)
	dc.DrawStringWrapped("UPPER MIDDLE", W/2, P, 0.5, 0, 0, 1.5, gg.AlignCenter)
	dc.DrawStringWrapped("LOWER MIDDLE", W/2, H-P, 0.5, 1, 0, 1.5, gg.AlignCenter)
	dc.DrawStringWrapped("LEFT MIDDLE", P, H/2, 0, 0.5, 0, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped("RIGHT MIDDLE", W-P, H/2, 1, 0.5, 0, 1.5, gg.AlignRight)
	if err := dc.LoadFontFace("/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf", 12); err != nil {
		panic(err)
	}
	dc.DrawStringWrapped(TEXT, W/2-P, H/2-P, 1, 1, W/3, 1.75, gg.AlignLeft)
	dc.DrawStringWrapped(TEXT, W/2+P, H/2-P, 0, 1, W/3, 2, gg.AlignLeft)
	dc.DrawStringWrapped(TEXT, W/2-P, H/2+P, 1, 0, W/3, 2.25, gg.AlignLeft)
	dc.DrawStringWrapped(TEXT, W/2+P, H/2+P, 0, 0, W/3, 2.5, gg.AlignLeft)
	dc.SavePNG(filename)
}
