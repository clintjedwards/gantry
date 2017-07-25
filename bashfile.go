package main

const bashrcContent = `
#Prompt settings
PS1="\[\e[34m[\$GANTRY_HOST]\] gantry $ \[\e[0m\]"

#Shortcuts
function deploy() {
	{
		docker-compose -f docker-compose.production.yml pull &&
		docker-compose -f docker-compose.production.yml down &&
		docker-compose -f docker-compose.production.yml up -d
	} || {
		echo "Command failed: Make sure you're in your project's correct docker directory"
		echo "Gantry expects compose file named: docker-compose.production.yml"
	}
}

function push() {
	{
		docker-compose -f docker-compose.production.yml pull &&
		docker-compose -f docker-compose.production.yml up -d
	} || {
		echo "Command failed: Make sure you're in your project's correct docker directory"
		echo "Gantry expects compose file named: docker-compose.production.yml"
	}
}

function help(){
	echo "Help:"
	echo "============"
	echo "gantry deploy - Pull, full shutdown, and restart of environment"
	echo "gantry push   - Pull and quick restart of environment"
}

function gantry(){
	case "$1" in
	"deploy")
		deploy
		;;
	"push")
		push
		;;
	*)
		help
		;;
	esac
}
`
