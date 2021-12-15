echo "choose yes or no!!!"

while true; do
  read -r -p "yes or no: " yn
  case ${yn} in
    [Yy]* ) break ;;
    [Nn]* ) exit 1 ;;
  esac
done

echo "some process..."
