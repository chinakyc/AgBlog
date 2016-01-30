angular.module("blogApp")
    .controller("HomeCtrl", function ($scope) {
        $scope.drawSkillChart = function(skill) {
            var doughnutData = [
              {
                value: skill.value,
                color:"#74cfae"
              },
              {
                value : 100 - skill.value,
                color : "#3c3c3c"
              }
            ];
            var myDoughnut = new Chart(document.getElementById(skill.name).getContext("2d")).Doughnut(doughnutData);
        }
    });
