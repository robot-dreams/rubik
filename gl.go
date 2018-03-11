package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

const (
	// These values must match the declarations in the vertex shader.
	perspectiveUniform = "perspective"
	viewUniform        = "view"
)

func initOpenGL() error {
	if err := gl.Init(); err != nil {
		return err
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	gl.Enable(gl.DEPTH_TEST)
	// gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)

	return nil
}

func loadShaderSource(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	data = append(data, '\x00')
	return string(data), nil
}

func checkShaderError(source string, shader uint32) error {
	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))
		return fmt.Errorf("Failed to compile %v: %v", source, log)
	}
	return nil
}

func newShader(path string, shaderType uint32) (uint32, error) {
	source, err := loadShaderSource(path)
	if err != nil {
		return 0, err
	}
	csources, free := gl.Strs(source)
	defer free()

	shader := gl.CreateShader(shaderType)
	gl.ShaderSource(shader, 1, csources, nil)
	gl.CompileShader(shader)

	err = checkShaderError(source, shader)
	if err != nil {
		return 0, err
	}
	return shader, nil
}

func checkProgramError(program uint32) error {
	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))
		return fmt.Errorf("Failed to link program: %v", log)
	}
	return nil
}

func newProgram(
	vertexShaderPath string,
	fragmentShaderPath string,
) (uint32, error) {
	vertexShader, err := newShader(vertexShaderPath, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}
	defer gl.DeleteShader(vertexShader)

	fragmentShader, err := newShader(fragmentShaderPath, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}
	defer gl.DeleteShader(fragmentShader)

	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	err = checkProgramError(program)
	if err != nil {
		return 0, err
	}
	return program, nil
}

func setUniformMatrix4fv(program uint32, name string, m mgl32.Mat4) {
	ul := gl.GetUniformLocation(program, gl.Str(name+"\x00"))
	if ul == -1 {
		panic(fmt.Errorf("Could not find uniform %v", name))
	}
	gl.ProgramUniformMatrix4fv(program, ul, 1, false, &m[0])
}
