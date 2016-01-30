angular.module("blogApp").service("ArticleService", ["$http", "API_URLS", "AuthService", "USER_ROLES", "BASE_CONFIG",
    function($http, API_URLS, AuthService, USER_ROLES, BASE_CONFIG) {

        this.get = function(articleID) {
            return $http.get(API_URLS.article + "/" + articleID);
        };

        // delete article real authorization in server-side
        this.deleteArticle = function(articleID) {
            if (AuthService.isAuthorized(USER_ROLES.admin)) {
                return $http.delete(API_URLS.article + "/" + articleID);
            }
        };

        // add article real authorization in server-side
        this.add = function(article) {
            if (AuthService.isAuthorized([USER_ROLES.editor, USER_ROLES.admin])) {
                return $http.post(API_URLS.article, article)
            }
        };

        // edit article real authorization in server-side
        this.edit = function(article) {
            if (AuthService.isAuthorized(USER_ROLES.admin)) {
                return $http.put(API_URLS.article + "/" + article.id, article)
            }
        };

        this.all = function(categoryId, page) {
            // page_num start from 0
            page -= 1

            return $http.get(API_URLS.articles + categoryId, {
                params: { limit: BASE_CONFIG.blogListPerPage , offset: page * BASE_CONFIG.blogListPerPage}
            })
        };

        return this
    },
]);
