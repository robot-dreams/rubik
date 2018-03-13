#version 410 core

layout (location = 0) in vec3 position;
layout (location = 1) in vec3 color;
layout (location = 2) in vec3 normal;

uniform mat4 view;
uniform mat4 perspective;

out vec3 fragmentPosition;
out vec3 stickerColor;
out vec3 stickerNormal;

void main() {
    gl_Position = perspective * view * vec4(position, 1.0);
    fragmentPosition = position;
    stickerColor = color;
    stickerNormal = normal;
}
