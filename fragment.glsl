#version 410 core

in vec3 stickerColor;

out vec4 fragmentColor;

void main() {
    fragmentColor = vec4(stickerColor, 1.0);
}
