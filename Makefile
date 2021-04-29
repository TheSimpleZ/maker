PHONY: targetx
targetx:
	echo "Hello there"
	exit 1
	echo "should not make it"



## This is a special rule that enables the use of
## $(filter-out $@,$(MAKECMDGOALS))
## '%:' is a rule that matches anything
## '@:' is a recipe that does nothing
## the @ means 'do it silently'
## in short: If recipe can't be found, do nothing and do it silently
%:
	@:
