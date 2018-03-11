package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
	"unsafe"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

func shaderSource(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	data = append(data, '\x00')
	return string(data), nil
}

func checkShaderErr(source string, shader uint32) error {
	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))
		return fmt.Errorf("failed to compile %v: %v", source, log)
	}
	return nil
}

func compileShader(path string, shaderType uint32) (uint32, error) {
	source, err := shaderSource(path)
	if err != nil {
		return 0, err
	}

	csources, free := gl.Strs(source)
	defer free()

	shader := gl.CreateShader(shaderType)
	gl.ShaderSource(shader, 1, csources, nil)
	gl.CompileShader(shader)

	err = checkShaderErr(source, shader)
	if err != nil {
		return 0, err
	}
	return shader, nil
}

func checkProgramErr(program uint32) error {
	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))
		return fmt.Errorf("failed to link program: %v", log)
	}
	return nil
}

func glProgram(
	vertexShaderPath string,
	fragmentShaderPath string,
) (uint32, error) {
	vertexShader, err := compileShader(vertexShaderPath, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}
	defer gl.DeleteShader(vertexShader)

	fragmentShader, err := compileShader(fragmentShaderPath, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}
	defer gl.DeleteShader(fragmentShader)

	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	err = checkProgramErr(program)
	if err != nil {
		return 0, err
	}
	return program, nil
}

func mustSetUniformMatrix4fv(program uint32, name string, m mgl32.Mat4) {
	ul := gl.GetUniformLocation(program, gl.Str(name+"\x00"))
	if ul == -1 {
		panic(fmt.Errorf("uniform %v not found", name))
	}
	gl.ProgramUniformMatrix4fv(program, ul, 1, false, &m[0])
}

func glObjects() (vao, vbo, ebo uint32) {
	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)
	gl.GenBuffers(1, &ebo)
	return vao, vbo, ebo
}

func glDraw(program, vao, vbo, ebo uint32, data []float32, indices []uint32) {
	gl.UseProgram(program)
	gl.BindVertexArray(vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(data), gl.Ptr(data), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*len(indices), gl.Ptr(indices), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 24, nil)
	gl.EnableVertexAttribArray(0)

	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 24, unsafe.Pointer(uintptr(12)))
	gl.EnableVertexAttribArray(1)

	gl.DrawElements(gl.TRIANGLES, int32(len(indices)), gl.UNSIGNED_INT, nil)
}

func glRadians(degrees float64) float32 {
	return float32(degrees * math.Pi / 180)
}
