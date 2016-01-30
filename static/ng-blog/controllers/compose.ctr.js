angular.module("blogApp")
    .controller("ComposeCtrl", ["$scope", "$state", "MESSAGE_TYPE", "ArticleService", "Session", "article", function ($scope, $state, MESSAGE_TYPE, ArticleService, Session, article) {

            if (article) {
                $scope.article = article
            } else {
                $scope.Compose = true
                $scope.article = {
                    title: "",
                    markdown: "",
                    content: "",
                    category: "",
                };
            }

            $scope.error = null;
            $scope.submit = function(article) {
                    article.markdown = Editor.getMarkdown();
                    article.content = Editor.getHTML();

                    if ($state.current.name == "compose") {
                        ArticleService.add(article).success(function(response) {
                            $state.transitionTo('articleDeteail', {id: response.id})
                            Session.flashMessage("添加成功", MESSAGE_TYPE.success)
                        }).error(function(reponse) {
                            $scope.error = reponse.error;
                        })
                    } else {
                        ArticleService.edit(article).success(function(response) {
                            $state.transitionTo('articleDeteail', {id: response.id})
                            Session.flashMessage("修改成功", MESSAGE_TYPE.success)
                        }).error(function(reponse) {
                            $scope.error = reponse.error;
                        })
                    }
                }

            var Editor = editormd("editormd", {
                    height : 640,
                    path   : "static/js/editormd/lib/",
                    saveHTMLToTextarea : true,
                    codeFold : true,
            });
        }
    ]);
