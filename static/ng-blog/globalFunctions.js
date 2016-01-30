angular.module("blogApp").run(function ($rootScope, $window, $state, AuthService, AUTH_EVENS, Session, BASE_CONFIG) {

        $rootScope.baseConfig = BASE_CONFIG

        if (!AuthService.isAuthenticated() && $window.sessionStorage["userInfo"]) {
            var credentials = JSON.parse($window.sessionStorage["userInfo"]);
            AuthService._login(credentials);
        }

        $rootScope.$on("$stateChangeStart", function (event, next, toParams) {
            var authorizedRoles = next.data.authorizedRoles;

            Session.setPreviousState($state.current.name, $state.params)

            if (authorizedRoles && !AuthService.isAuthorized(authorizedRoles)) {
                event.preventDefault();
                if (AuthService.isAuthenticated()) {
                    // user is not allowed
                    $rootScope.$broadcast(AUTH_EVENS.notAuthorized);
                } else {
                    // user is not login
                    $rootScope.$broadcast(AUTH_EVENS.notAuthenticated);
                }
            }
        });
    })
