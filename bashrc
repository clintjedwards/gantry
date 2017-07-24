#Prompt settings
PS1="\[\e[34m[\$GANTRY_HOST]\] gantry $ \[\e[0m\]"

#Shortcuts
function full_deploy() {
    docker-compose -f docker-compose.producution.yml pull
    docker-compose -f docker-compose.production.yml down
    docker-compose -f docker-compose.production.yml up -d
}

function deploy() {
    docker-compose -f docker-compose.production.yml pull
    docker-compose -f docker-compose.production.yml up -d
}