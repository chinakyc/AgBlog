angular.module("blogApp")
    .controller("ArticleCtrl", ["$scope", "$sce", "article", "ArticleService", "AuthService", "Session",
        function ($scope, $sce, article, ArticleService, AuthService, Session) {
            article.content = $sce.trustAsHtml(article.content)

            $scope.article = article

            $scope.deleteArticle = function(articleId) {
                ArticleService.deleteArticle(articleId).success(function (data) {
                    Session.flashMessage("删除成功")
                    $scope.goBack()
                })
            }
        }
    ]);
