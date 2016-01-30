// https://medium.com/opinionated-angularjs/techniques-for-authentication-in-angularjs-applications-7bbf0346acec#.nno3rovhb
// check `response.statue` after each $http request 
angular.module("blogApp")
.factory("AuthInterceptor", ["$rootScope", "$q", "Session", "AUTH_EVENS", 
    function($rootScope, $q, Session, AUTH_EVENS) {
        return {
            responseError: function(response) {
                $rootScope.$broadcast({
                    401: AUTH_EVENS.notAuthenticated,
                    403: AUTH_EVENS.notAuthorized,
                    419: AUTH_EVENS.sessionTimeout,
                    440: AUTH_EVENS.sessionTimeout
                }[response.status], response);
                return $q.reject(response);
            }
        };
    }
]);
