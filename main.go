//main
//Author: Neil Balaskandarajah
//Created on: 09/21/2019
//Main file that handles all of the graphics and the model of the arm

package main

import (
	// "fmt"
	"github.com/h8gi/canvas"
	"golang.org/x/image/colornames"
	// "math"
	// "time"
)

//Constants

const width int = 1920      //WIDTH is the width of the window
const height int = 1080     //HEIGHT is the height of the window
const fps int = 100         //FPS is the frame rate of the animation
const fontSize float64 = 60 //FONT_SIZE is the font size for the canvas

//variables
var c canvas.Canvas //canvas instance
var robotArm Arm    //arm struct

//create the arm struct to be used and run the graphics
func main() {
	//create a new canvas instance
	c := canvas.NewCanvas(&canvas.CanvasConfig{
		Width:     width,
		Height:    height,
		FrameRate: fps,
		Title:     "Arm Simulator",
	})

	//set up the canvas
	c.Setup(func(ctx *canvas.Context) {
		setUpCanvas(ctx)
	})

	//create the arm
	createArm()

	//create the arm
	c.Draw(func(ctx *canvas.Context) {
		//update the arm
		updateModel()

		//clear the canvas
		ctx.SetColor(colornames.Black) //set the bg color
		ctx.Clear()                    //empty the canvas

		//save canvas state
		ctx.Push()

		drawRobot(ctx) //draw the robot to the screen

		//display the data to the screen
		displayData(ctx)

		//restore canvas state
		ctx.Pop()
	})

} //end main

//create the arm
func createArm() {
	//set the values for the arm
	startPt := point{float64(width / 2), 0}
	length := 24.0
	angle := ToRadians(0)

	//create the arm
	robotArm = Arm{
		start:    startPt,
		length:   length,
		angle:    angle,
		topSpeed: 240} //60 degrees per second, 10RPM

	//create the PID controller
	pid := pidcontroller{
		kP: 2,
		kI: 0.002,
		kD: 0.2}

	//set the PID controller for the arm
	robotArm.setPIDController(pid)
} //end createArm

//Update the arm for drawing purposes
func updateModel() {
	//move with PID control until the target is reached
	robotArm.movePID(160, robotArm.getAngleDeg(), 1)
} //end updateModel

//SIMULATOR

//draw the robot to the display
//ctx *canvas.Context - responsible for drawing
func drawRobot(ctx *canvas.Context) {
	//switch to the arm color
	colors := robotArm.getColor()
	ctx.SetRGB255(colors[0], colors[1], colors[2])

	//draw the robot arm as lines between the joint points
	ctx.DrawLine(robotArm.start.x, robotArm.start.y,
		robotArm.getEndPtPxl().x, robotArm.getEndPtPxl().y)

	ctx.Stroke()
} //end drawRobot

//display the parameters of the robot onto the screen
//ctx *canvas.Context - responsible for drawing
func displayData(ctx *canvas.Context) {
	ctx.SetColor(colornames.White)

	//display the start and end coords
	displayPointCoords(ctx, robotArm.getStartPtIn(), 1400, fontSize)
	displayPointCoords(ctx, robotArm.getEndPtIn(), 1400, 70+fontSize)

	//display the angle of the arm (combine with point and make into helper function)
	drawFloat(ctx, robotArm.angle, 1400, 140+fontSize)
	drawFloat(ctx, robotArm.getAngleDeg(), 1600, 140+fontSize)
} //end displayData
