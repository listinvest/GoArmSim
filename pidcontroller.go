//pidcontroller
//Author: Neil Balaskandarajah
//Created on: 09/25/2019
//A simple PID controller that assumes regular loop intervals

package main

import (
	"math"
)

//pid controller struct with the three gains
type pidcontroller struct {
	//configured attributes
	kP float64 //proportionality constant
	kI float64 //integral constant
	kD float64 //derivative constant

	//calculated attributes
	errorSum  float64 //sum of all errors
	lastError float64 //last error for derivative calculation
	epsilon   float64 //the range to be in to be considered "at goal"
	atTarget  bool    //whether within epsilon bounds

	goal float64 //goal controller is trying to reach
} //end struct

//calculate the PID output based on the setpoint, current value and tolerance
//float64 setpoint - the desired goal value
//float64 current - the current value
//float64 epsilon - the range you can be in to be considered "at goal"
func (pid *pidcontroller) calcPID(setpoint, current, epsilon float64) float64 {
	pid.goal = setpoint
	//get the error
	error := pid.goal - current

	//update atTarget
	pid.atTarget = math.Abs(error) <= epsilon

	//P value
	pOut := pid.kP * error //output proportional to error

	//I value
	pid.errorSum += error //add onto the error sum
	iOut := pid.kI * pid.errorSum

	//D value
	dError := (error - pid.lastError)
	dOut := pid.kD * dError

	return pOut + iOut + dOut //sum of outputs
} //end calcPID

//calculate the voltage required to hold an arm up at a certain angle
//Arm a - arm to hold up
func calcFFArm(a *Arm) float64 {
	return (a.calcGravTorque() * a.motor.kResistance) / (a.kT * a.gearRatio)
	// return 0
} //end calcFFArm

//return whether the error is within the epsilon bounds or not
func (pid pidcontroller) isDone() bool {
	return pid.atTarget
} //end isDone
