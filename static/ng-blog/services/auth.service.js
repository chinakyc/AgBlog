// https://medium.com/opinionated-angularjs/techniques-for-authentication-in-angularjs-applications-7bbf0346acec#.nno3rovhb
// the login related to authentication and authorization
angular.module("blogApp")
    .factory("AuthService", [ "$http", "$window", "$rootScope", "Session", "AUTH_EVENS", "API_URLS", 
        function($http, $window,$rootScope, Session, AUTH_EVENS, API_URLS) {
            var authService = {}

            // the login function
            authService.login = function(user) {
                return $http.post(API_URLS.login, user).success(function(loginData) {
                    authService._login(loginData)
                })
            };

            // call this to load user when login success
            authService._login = function(loginData) {
               // set the browser session, to avoid relogin on refresh
               $window.sessionStorage["userInfo"] = JSON.stringify(loginData);

               // update current user into the Session service
               Session.createUser(loginData.token, loginData.role, loginData.user);

               // Set token to headers
               $http.defaults.headers.common.Authorization = 'Bearer ' + Session.userToken;

               //fire event of successful login
               $rootScope.$broadcast(AUTH_EVENS.loginSuccess);

            }

            authService.logOut = function () {
                Session.destroy()
                $http.defaults.headers.common.Authorization = "";  
                $window.sessionStorage.removeItem("userInfo");
                $rootScope.$broadcast(AUTH_EVENS.logoutSuccess);
            }

            // check user authorized to access
            authService.isAuthenticated = function() {
                return !!Session.userToken;
            }

            // check user authorized to access
            authService.isAuthorized =  function(authorizedRoles) {
                if (!angular.isArray(authorizedRoles)) {
                    authorizedRoles = [authorizedRoles];
                }
                return (authService.isAuthenticated() &&
                        authorizedRoles.indexOf(Session.userRole) !== -1);
            };
            
            return authService;
        }
    ]);
