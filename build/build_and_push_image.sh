GOOS=linux make push
for tag in ${tags//,/ }
do
  docker tag $docker_repo:local $docker_repo:$tag
  docker push $docker_repo:$tag
done
