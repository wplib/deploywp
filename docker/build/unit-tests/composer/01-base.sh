#!/bin/bash
# Created on 2020-04-27T15:46:35+1000, using template:01-base.sh.tmpl and json:gearbox.json

p_info "pantheon" "Release test started."

if id -u gearbox
then
	c_ok "Gearbox user found."
else
	c_err "Gearbox user NOT found."
fi

if id -g gearbox
then
	c_ok "Gearbox group found."
else
	c_err "Gearbox group NOT found."
fi

p_info "pantheon" "Release test finished."

