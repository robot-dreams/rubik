#version 410 core

uniform vec3 viewPosition;

in vec3 fragmentPosition;
in vec3 stickerColor;
in vec3 stickerNormal;

out vec4 fragmentColor;

void main() {
    float ambientStrength = 0.15;
    float diffuseStrength = 0.7;
    float specularStrength = 0.9;
    float specularPower = 16;

    vec3 lightPosition = viewPosition;
    vec3 lightColor = vec3(1.0);

    vec3 ambient = ambientStrength * lightColor;

    vec3 lightDirection = normalize(lightPosition - fragmentPosition);
    float diffuseAmount = max(dot(stickerNormal, lightDirection), 0.0);
    vec3 diffuse = diffuseStrength * diffuseAmount * lightColor;

    vec3 viewDirection = normalize(viewPosition - fragmentPosition);
    vec3 reflectDirection = reflect(-lightDirection, stickerNormal);
    float specularAmount = pow(
            max(dot(viewDirection, reflectDirection), 0.0),
            specularPower);
    vec3 specular = specularStrength * specularAmount * lightColor;

    vec3 result = (ambient + diffuse + specular) * stickerColor;
    fragmentColor = vec4(result, 1.0);
}
