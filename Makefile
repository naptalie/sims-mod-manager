SHELL=/bin/bash

install_ubuntu_requirements:
	sudo apt-get update
	sudo apt-get install python3-pip
	sudo apt install python3.10-venv

source_venv:
	source .venv/bin/activate

install_python_requirements:
	pip install -r requirements.txt