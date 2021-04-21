state=$(turandot service list -o json | jq '.[0].instantiationState')

if [ $state != '"Instantiated"' ]
then
   state=$(turandot service list -o json | jq '.[0].instantiationState')
   echo $state
   sleep 2
fi



