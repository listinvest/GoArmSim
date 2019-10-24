//arm2
//Author: Neil Balaskandarajah
//Created on: 10/06/2019
//A two degree of freedom arm

package main

import (
	"math"
	"time"
)

//Arm2 is a two degree of freedom arm is made of two Arm structs, with the elbow dynamically changing start position
type Arm2 struct {
	arm1  *Arm       //the base joint (shoulder)
	arm2  *Arm       //the second joint (elbow)
	timer time.Timer //timer to delay arm tracking goal points
} //end struct

//Updates the position of the arms, translating the second joint start to the first joint end
func (a2 *Arm2) update() {
	a2.arm2.setStartPt(a2.arm1.getEndPtPxl())
} //end update

//updates the individual arms with zero voltage
func (a2 Arm2) rest() {
	a2.update()
	a2.arm1.update()
	a2.arm2.update()
} //end rest

//Set the color for both arms
//[3]int color - RGB values for the slice
func (a2 *Arm2) setArmColors(color [3]int) {
	a2.arm1.color = color
	a2.arm2.color = color
} //end setArmColors

//Check if the arm is stopped
//return - whether both joints are stopped
func (a2 Arm2) isStopped() bool {
	return a2.arm1.stopped && a2.arm2.stopped
} //end isStopped

//InverseKinematics calculates the joint angles given an endpoint
//Point p - endpoint in Cartesian space
//float64 ang1 - current angle of first joint
//float64 ang2 - current angle of second joint
//float64 a1 - length of first joint
//float64 a2 - length of second joint
//return - new first and second joint angles
func InverseKinematics(p Point, ang1, ang2, a1, a2 float64) (float64, float64) {
	r := PointDistance(Point{0, 0}, p) //distance from origin to point

	q2a := math.Acos((r*r - a1*a1 - a2*a2) / (2 * a1 * a2))                                     //second joint angle
	q1a := math.Atan2(p.y, p.x) - math.Abs(math.Atan((a2*math.Sin(q2a))/(a1+a2*math.Cos(q2a)))) //first joint angle

	q2b := -math.Acos((r*r - a1*a1 - a2*a2) / (2 * a1 * a2))                                    //second joint angle
	q1b := math.Atan2(p.y, p.x) + math.Abs(math.Atan((a2*math.Sin(q2a))/(a1+a2*math.Cos(q2a)))) //first joint angle

	//determine based on distance travelled
	if (math.Abs(q1a-ang1) + math.Abs(q2a-ang2)) < (math.Abs(q1b-ang1) + math.Abs(q2b-ang2)) {
		return q1a, q2a
	} //if
	return q1b, q2b
} //end moveToPoint
