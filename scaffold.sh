function main() {
  local NAMESPACE=$(echo $1 | sed '/\/$/! s|$|/|')
  local CHALLENGE=$2
  local LAST_FILE=$(ls $NAMESPACE | sort -V | tail -n 1)
  local LAST_COUNT=${LAST_FILE:0:1}
  local NEXT_COUNT=$((LAST_COUNT+1))
  local DIRECTORY=$NAMESPACE${NEXT_COUNT}_$CHALLENGE
  local SCRIPT="$DIRECTORY/$CHALLENGE.go"

  mkdir $DIRECTORY
  touch $SCRIPT
  cat << EOF >> $SCRIPT
package $CHALLENGE

func Challenge() {}
EOF
}

main $*
