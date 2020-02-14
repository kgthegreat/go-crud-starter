#! /bin/bash
entity=$1
entityUp="$(tr '[:lower:]' '[:upper:]' <<< ${entity:0:1})${entity:1}"
echo "Creating controller.."
cp controllers/topic.go controllers/${entity}.go
sed -i'.original' "s/topic/$entity/g" controllers/${entity}.go
sed -i'.original' "s/Topic/$entityUp/g" controllers/${entity}.go
echo "Creating model.."
cp models/topic.go models/${entity}.go
sed -i'.original' "s/topic/$entity/g" models/${entity}.go
sed -i'.original' "s/Topic/$entityUp/g" models/${entity}.go
echo "Creating repositories.."
cp repositories/topic.go repositories/${entity}.go
sed -i'.original' "s/topic/$entity/g" repositories/${entity}.go
sed -i'.original' "s/Topic/$entityUp/g" repositories/${entity}.go
echo "Creating views.."
cp -R templates/topic templates/${entity}
