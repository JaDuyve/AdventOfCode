day = $(shell date +'%-d')

new:
	@echo "Creating new file structure for day" $(day)"..."

	@if [ $(day) -lt 10 ] ; then \
  		mkdir calendar/day0$(day); \
  		cp template calendar/day0$(day)/part1.go; \
  		cp template calendar/day0$(day)/part2.go; \
  		touch calendar/day0$(day)/input; \
  	else \
  		mkdir calendar/day$(day); \
		cp template calendar/day$(day)/part1.go; \
		cp template calendar/day$(day)/part2.go; \
		touch calendar/day$(day)/input; \
    fi

	@git add calendar/

