class Grb
  GIT    = ENV['GRB_GIT']    || 'git'
  ORIGIN = ENV['GRB_ORIGIN'] || 'origin'

  COMMANDS = {
    :new     => {
      :commands => [
        '"#{GIT} push #{origin} #{current_branch}:refs/heads/#{branch}"',
        '"#{GIT} fetch #{origin}"',
        '"#{GIT} branch --track #{branch} #{origin}/#{branch}"',
        '"#{GIT} checkout #{branch}"'
      ]
    },

    :push     => {
      :commands => [
        '"#{GIT} push #{origin} #{branch}:refs/heads/#{branch}"',
        '"#{GIT} fetch #{origin}"',
        '"#{GIT} config branch.#{branch}.remote #{origin}"',
        '"#{GIT} config branch.#{branch}.merge refs/heads/#{branch}"',
        '"#{GIT} checkout #{branch}"'
      ]
    },

    :mv     => {
      :commands => [
        ' if(branch != branch_)
           "#{GIT} push #{origin} #{branch}:refs/heads/#{branch_}
            #{GIT} fetch #{origin}
            #{GIT} branch --track #{branch_} #{origin}/#{branch_}
            #{GIT} checkout #{branch_}
            #{GIT} branch -d #{branch}
            #{GIT} push #{origin} :refs/heads/#{branch}"
          else
           "#{GIT} push #{origin} #{current_branch}:refs/heads/#{branch}
            #{GIT} fetch #{origin}
            #{GIT} branch --track #{branch} #{origin}/#{branch}
            #{GIT} checkout #{branch}
            #{GIT} push #{origin} :refs/heads/#{current_branch}
            #{GIT} branch -d #{current_branch}"
          end'
      ]
    },

    :rm     => {
      :commands => [
        '"#{GIT} push #{origin} :refs/heads/#{branch}"',
        '"#{GIT} checkout master" if current_branch == branch',
        '"#{GIT} branch -d #{branch}"'
      ]
    },

    :pull      => {
      :commands => [
        '"#{GIT} fetch #{origin}"',
        'if local_branches.include?(branch) 
          "#{GIT} config branch.#{branch}.remote #{origin}\n" +
          "#{GIT} config branch.#{branch}.merge refs/heads/#{branch}"
        else
          "#{GIT} branch --track #{branch} #{origin}/#{branch}"
        end'
      ]
    }
  }

  def self.parse(opt)
    if COMMANDS.keys.include?(opt[:command].to_sym)
      current_branch,branch,branch_,origin = get_current_branch,opt[:branch],opt[:branch_],ORIGIN

      COMMANDS[opt[:command].to_sym][:commands].map {|x| exec_cmd(eval(x))}
    else
      puts opt[:command] + ' FALSE'
    end
  end

  def self.exec_cmd(str)
    return true unless str
    puts("\e[031m" + str + "\e[0m")
    system("#{str}")
  end

  def self.help(*args)
    #TODO
  end

  def self.get_current_branch
    (`git branch 2> /dev/null | grep '^\*'`).gsub(/\W/,'')
  end

  def self.local_branches
   (`git branch -l`).split(/\n/).map{|x| x.gsub(/\W/,'')}
  end
end
