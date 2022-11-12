#!/bin/bash
function gitInit(){
    git init
    echo "----------------------------------------------------"
    IFS= read -r -p "c or 3? " forwhom

    em=''
    un=''

    if [[ $forwhom = 'c' ]];then
        em="thir19426.cs@rmkec.ac.in"
        un="thrishalsundar"
    elif [[ $forwhom = '3' ]];then
        em="3shcodes@gmail.com"
        un="3shcodes"
    else
        echo "Poda mireeh"
        exit
    fi


    echo "----------------------------------------------------"
    git config --local user.name $em
    git config --local user.email $un
    echo "----------------------------------------------------"
    IFS= read -r -p "Enter RepoName: " repoName
    url="http://github.com/$un/$repoName.git"
    git remote add origin $url
    echo "Success: remote added as $url"
    exit 0
}

function gitPush(){
    git add .
    echo "----------------------------------------------------"
    git status
    echo "----------------------------------------------------"
    #echo "Commit msg: "
    IFS= read -r -p "Commit msg: " commitMsg
    echo "----------------------------------------------------"
    git commit -m "$commitMsg"
    echo "----------------------------------------------------"
    git push -u origin master
    exit 0
}

if [[ $1 = 'i' ]]; then 
    gitInit
elif [[ $1 = 'p' ]]; then
    gitPush
else 
    echo "Poda punda"
    exit
fi
