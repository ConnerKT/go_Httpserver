# Copy folder
rsync -r --exclude .git ./ ../go_httpserver

cd ../go_httpserver

# git a .
git add .

git commit -m "auto commit" 

git push origin main

