angular.module("blogApp").config(["$stateProvider", "$urlRouterProvider",  "USER_ROLES", "API_URLS", 
    function($stateProvider, $urlRouterProvider, USER_ROLES, API_URLS) {
        // unmatched url
        $urlRouterProvider.otherwise("/");

        $stateProvider.state("home", {
            url: "/",
            templateUrl: "templates/home.html",
            controller: "HomeCtrl",
            data: {
                authorizedRoles: null,
            },
        }).state("login", {
            url: "/login",
            templateUrl: "templates/login.html",
            controller: "LoginCtrl",
            data: {
                authorizedRoles: null,
            }
        }).state("about", {
            url: "/about",
            templateUrl: "templates/about.html",
            data: {
                authorizedRoles: null
            },
        }).state("blog", {
            url: "/blog",
            templateUrl: "templates/blog.html",
            controller: "BlogCtrl",
            data: {
                authorizedRoles: null
            },
            resolve: {
                categorys: function(CategoryService) {
                    return CategoryService.get().then(function(response) {
                        return response.data;
                    })
                }
            }
        }).state("blog.articles", {
            url: "/{category:int}/{page:int}",
            templateUrl: "templates/blog.list.html",
            controller: "ArticleListCtrl",
            data: {
                authorizedRoles: null
            },
            resolve: {
                articles: function(ArticleService, $stateParams) {
                    return ArticleService.all($stateParams.category, $stateParams.page).then(function(response) {
                        return response.data;
                    })
                }
            }
        }).state("articleDeteail", {
            url: "/article/{id:int}",
            templateUrl: "templates/article.html",
            controller: "ArticleCtrl",
            data: {
                authorizedRoles: null
            },
            resolve: {
                article: function (ArticleService, $stateParams) {
                    return ArticleService.get($stateParams.id).then(function(response) {
                        return response.data;
                    })
                }
            }
        }).state("compose", {
            url: "/compose",
            templateUrl: "templates/compose.html",
            controller: "ComposeCtrl",
            data: {
                authorizedRoles: [USER_ROLES.admin, USER_ROLES.editor]
            },
            resolve: {
                article: function($q) {
                    var deferred = $q.defer();
                    deferred.resolve(null);
                    return deferred.promise;
                }
            }
        }).state("edit", {
            url: "/edit/:id",
            templateUrl: "templates/compose.html",
            controller: "ComposeCtrl",
            data: {
                authorizedRoles: [USER_ROLES.admin]
            },
            resolve: {
                article: function (ArticleService, $stateParams) {
                    return ArticleService.get($stateParams.id).then(function(response) {
                        return response.data;
                    })
                }
            }
        })
    }
]);
