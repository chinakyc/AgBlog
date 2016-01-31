// main application module
var blogApp = angular.module("blogApp", ["ui.router", "angular-loading-bar", "ngAnimate"])
// Constants use for login 
.constant("USER_ROLES", {
    guest: 2,
    editor: 1,
    admin: 0,
}).constant("AUTH_EVENS", {
    loginSuccess: "auth-login-success",
    loginFailed: "auth-login-failed",
    logoutSuccess: "auth-logout-success",
    sessionTimeout: "auth-session-timeout",
    notAuthenticated: "auth-not-authenticated",
    notAuthorized: "auth-not-authorized",
// Constants API URL 
}).constant("API_URLS", {
    login: "/login",
    categorys: "/categorys",
    owner: "/owner",
    article: "/article",
    articles: "/category/",
// Constants message
}).constant("MESSAGE_TYPE", {
    error: "error",
    success: "success",
}).constant("MESSAGE_EVEN", {
    flashed: "message-flashed"
}).constant("BASE_CONFIG", {
    blogTitle: "Azul's blog 做一个靠谱的人",
    blogListPerPage: 8
})
// Adding the auth interceptor here
.config(function ($httpProvider) {
    $httpProvider.interceptors.push([
        "$injector",
        function ($injector) {
            return $injector.get("AuthInterceptor");
       }
    ]);
});
