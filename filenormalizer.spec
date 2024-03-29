%define debug_package   %{nil}
%define _build_id_links none
%define _name   filenormalizer
%define _prefix /opt
%define _version 1.001
%define _rel 0
%define _arch x86_64
%define _binaryname fnormalizer

Name:       filenormalizer
Version:    %{_version}
Release:    %{_rel}
Summary:    filenormalizer

Group:      Administration tools
License:    GPL2.0
URL:        https://github.com/jeanfrancoisgratton/filenormalizer

Source0:    %{name}-%{_version}.tar.gz
BuildArchitectures: x86_64
#BuildRequires: libvirt-devel,wget,gcc
#Requires: libvirt-devel,libvirt,virt-clone,sudo,postgresql-contrib


%description
Filesystem utilities

%prep
#%setup -q
%autosetup

%build
cd %{_sourcedir}/%{_name}-%{_version}/src
PATH=$PATH:/opt/go/bin go build -o %{_sourcedir}/%{_binaryname} .
strip %{_sourcedir}/%{_binaryname}

%clean
rm -rf $RPM_BUILD_ROOT

%pre

%install
#%{__mkdir_p} "$RPM_BUILD_ROOT%{_prefix}/bin"
#install -Dpm 0755 %{buildroot}/%{_name} "$RPM_BUILD_ROOT%{_prefix}/bin/"
install -Dpm 0755 %{_sourcedir}/%{_binaryname} %{buildroot}%{_bindir}/%{_binaryname}

%post
strip %{_prefix}/bin/%{_binaryname}

%preun

%postun

%files
%defattr(-,root,root,-)
%{_bindir}/%{_binaryname}


%changelog
* Thu Aug 03 2023 builder <builder@famillegratton.net> 1.001-0
- Fixed issue where no flags are passed (jean-francois@famillegratton.net)
- Fixed shell issue for NAS (jean-francois@famillegratton.net)

* Thu Aug 03 2023 builder <builder@famillegratton.net> 1.000-0
- Bumped to Go 1.20.6 and prod release (jean-francois@famillegratton.net)
- Bumped go pkgs (jean-francois@famillegratton.net)
- Perm change (jean-francois@famillegratton.net)
- updated go to 1.20.5 (jean-francois@famillegratton.net)


* Sun May 14 2023 builder <builder@famillegratton.net> 0.100-2
- new package built with tito

