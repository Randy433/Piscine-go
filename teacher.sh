InterviewNumber=$(head -n 179 streets/Buckingham_Place | tail -n 1 | cut -d "#" -f2)
echo $InterviewNumber
cat interviews/interview-"$InterviewNumber"
echo $MAIN_SUSPECT

