def main():
    user_list = list()
    with open("./user-data/userxl.txt","r") as file:
        content = file.readline()
        while content!="":
            email,handle= content.strip("\n").split('\t')
            user_list.append({"email":email,"handle":handle})
            content = file.readline()
    with open("./user-data/data.json","w") as file:
        file.write(str(user_list))
if __name__ =="__main__":
    main()
