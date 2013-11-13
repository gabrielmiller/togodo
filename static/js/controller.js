function togodo($scope){
    $scope.input = '';

    $scope.items = [
        'Write more golang apps',
        'Play with angular.js'
    ];

    $scope.addItem = function(){
        if($scope.input.length > 0){
            // call api endpoint to add item
            $scope.items.push($scope.input);
            $scope.input = '';
        }
    };

    $scope.removeItem = function(index){
        var deathToTask = confirm("Remove item '"+$scope.items[index]+"' from list?");
        if(deathToTask){
            // Call api endpoint to delete item
            $scope.items.splice(index,1);
        }
    };
}
