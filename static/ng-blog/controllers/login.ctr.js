angular.module("blogApp")
    .controller("LoginCtrl", ["$scope", "AuthService", "AUTH_EVENS", function ($scope, AuthService, AUTH_EVENS) {
        $scope.credentials = {
            email: "",
            password: ""
        };
        $scope.error = null;
        $scope.login = function(credentials) {
            AuthService.login(credentials).error(function(loginData) {
                $scope.error = loginData.error;
            })
        }
    }]);
