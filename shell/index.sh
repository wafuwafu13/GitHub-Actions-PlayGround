echo "choose yes or no!!!"

flag=${FLAG:="FALSE"}
if [ "${flag}" = "FALSE" ]; then
  while true; do
    read -r -p "yes or no: " yn
    case ${yn} in
      [Yy]* ) break ;;
      [Nn]* ) exit 1 ;;
    esac
  done
fi

echo "some process..."
