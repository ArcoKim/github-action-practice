#!/bin/bash
curl -O https://bootstrap.pypa.io/get-pip.py
python3 get-pip.py --user
cd /home/ec2-user
pip3 install -r requirements.txt