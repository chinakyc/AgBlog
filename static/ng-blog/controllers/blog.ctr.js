angular.module("blogApp")
    .controller("BlogCtrl", function ($scope, $state, categorys) {
        $scope.categorys = categorys
    });
