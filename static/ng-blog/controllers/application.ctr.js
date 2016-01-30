angular.module("blogApp")
    .controller("ApplicationController", ["$rootScope", "$scope", "$state", "AuthService", "AUTH_EVENS", "USER_ROLES", "$state", "Session", "MESSAGE_TYPE", "$http",
            function ($rootScope, $scope, $state, AuthService, AUTH_EVENS, USER_ROLES, $state, Session, MESSAGE_TYPE, $http) {

                $http.get("/owner").then(function (response) {
                    $scope.owner = response.data
                })

                $scope.userRoles = USER_ROLES;
                $scope.isAuthorized = AuthService.isAuthorized;
                $scope.isAuthenticated = AuthService.isAuthenticated;
                $scope.messageType = MESSAGE_TYPE;

                $scope.goBack = function () {
                    if (Session.previousState.state) {
                        $state.transitionTo(Session.previousState.state, Session.previousState.params)
                    } else {
                        $state.transitionTo("home")
                    }
                }

                $scope.logOut = AuthService.logOut;

                function notAuthorized () {
                    Session.flashMessage("未授权", MESSAGE_TYPE.error)
                }

                function notAuthenticated () {
                    Session.flashMessage("请登录", MESSAGE_TYPE.error)
                    $state.go("login")
                }

                function loginSuccess () {
                    Session.flashMessage("欢迎回来", MESSAGE_TYPE.success)
                    $state.go("home")
                }

                function logoutSuccess () {
                    Session.flashMessage("已登出", MESSAGE_TYPE.success)
                    $state.go("home")
                }

                $scope.$on(AUTH_EVENS.notAuthorized, notAuthorized);
                $scope.$on(AUTH_EVENS.notAuthenticated, notAuthenticated);
                $scope.$on(AUTH_EVENS.sessionTimeout, notAuthenticated);
                $scope.$on(AUTH_EVENS.logoutSuccess, logoutSuccess);
                $scope.$on(AUTH_EVENS.loginSuccess, loginSuccess);
            }
        ]
    );
