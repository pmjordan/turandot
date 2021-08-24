state=$(turandot service list -o json | jq '.[0].instantiationState')

while [ $state != '"Instantiated"' ]
do
   state=$(turandot service list -o json | jq '.[0].instantiationState')
   echo $state
   sleep 2
done



