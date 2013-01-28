

template node['mongodb']['configfile'] do
  source "mongodb.conf.erb"
  variables(
    :dbpath => node['mongodb']['dbpath'],
    :logpath => node['mongodb']['logpath'],
    :port => node['mongodb']['port'],
    :journal => node['mongodb']['journal'],
    :httpinterface => node['mongodb']['nohttpinterface'],
    :rest => node['mongodb']['rest'],
    :replication_set => node['mongodb']['replicaset'],
    :master => node['mongodb']['master'],
    :oplogsize => node['mongodb']['oplogsize'],
    :slave => node['mongodb']['slave'],
    :source => node['mongodb']['source']
  )
  owner "root"
  group "root"
  mode "0644"
  action :create
end

